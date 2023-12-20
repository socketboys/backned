import warnings
import sys
import os
from re import search
import typer
import scipy
import torch
from transformers import MBartForConditionalGeneration, MBart50TokenizerFast
from transformers import AutoProcessor, SeamlessM4Tv2Model
from faster_whisper import WhisperModel
from transformers import  AutoModelForSeq2SeqLM
from transformers import NllbTokenizerFast
from transformers.pipelines import pipeline
import ssl
from config import root_dir, languages,gender
from loguru import logger
from multiprocessing import Pool
import time
from itertools import repeat
# import ray
# from ray.util.multiprocessing import Pool


ssl._create_default_https_context = ssl._create_unverified_context

logger.add("./logs/{time}.log", level="TRACE", rotation="100 MB")

logger.info(f"Running on {sys.platform}")


class Pipeline:
    def __init__(self,input_path,audio_name,language,gender):
        logger.info("Class initialized")
        self.input_path = input_path
        self.audio_name = audio_name
        self.language = language
        self.gender = gender
        self.device = torch.device('cuda' if torch.cuda.is_available() else 'cpu')
        logger.info(f"Device: {self.device}")
        self.translated = list()

    
    
    def translate(self,text,lang):
        try:
            tokenizer = NllbTokenizerFast.from_pretrained("facebook/nllb-200-distilled-600M")
            model = AutoModelForSeq2SeqLM.from_pretrained("facebook/nllb-200-distilled-600M").to(self.device)
        except Exception as e:  
            logger.error(f"Error occurred while loading the model/tokenizer: {str(e)}")
            raise typer.Exit(1)
        try:
            logger.info(f"Translating in {lang}")
            inputs = tokenizer(text, return_tensors="pt", padding = True).to(self.device)

            translated_tokens = model.generate(
    **inputs, forced_bos_token_id=tokenizer.lang_code_to_id[lang]
)
            text = tokenizer.batch_decode(translated_tokens, skip_special_tokens=True)
            logger.info("Translation Done")
            return text[0]
        except Exception as e:
            logger.error(f"Error while translating text:{str(e)}")
            raise typer.Exit(1)
                    
    
    def translated_sub(self,file,lang):
        try:
            logger.info("Creating Subtitles")
            output_file = f"{root_dir}/subtitle/" + file + \
                "_" + languages[lang]["tts"][-3:-1] + ".srt"
            input_file = f"{root_dir}/input/" + file + ".srt"

            logger.info("Reaching out to translator function...")
            with open(input_file, 'r', encoding="utf-8") as infile, open(output_file, 'x+', encoding="utf-8") as outfile:
                for line in infile:

                    match = search(
                        r'\d+\n\d{2}:\d{2}:\d{2},\d{3} --> \d{2}:\d{2}:\d{2},\d{3}\n|\d+\n', line)
                    to_translate = search(r'^[a-zA-Z]', line)
                    if match:
                        outfile.write(line)
                    elif to_translate:

                        translated_text = self.translate(line, languages[lang]["translate"])
                        self.translated.append(translated_text)
                        outfile.write(translated_text + "\n\n")
            logger.info("Translated.")
            logger.info("Subtitles are created")
        except Exception as e:
            logger.error(f"Error while translating subtitles:{str(e)}")
            raise typer.Exit(1)
            
    def tts(self,file, text, lang,gender):
        try:
            logger.info("Starting Text to Speech")
            file_name = f"{root_dir}/audio/" + file + "_" + lang[-3:-1] + ".wav"
            try:
               processor = AutoProcessor.from_pretrained("facebook/seamless-m4t-v2-large")
               model = SeamlessM4Tv2Model.from_pretrained("facebook/seamless-m4t-v2-large").to(self.device)
            except Exception as e:
                logger.error(f"Error while loading TTS model:{str(e)}")
                raise typer.Exit(1)
            
            text_inputs = processor(text = text, src_lang=lang, return_tensors="pt").to(self.device)
            output = model.generate(**text_inputs, tgt_lang=lang, speaker_id=gender)[0].cpu().numpy().squeeze()

            scipy.io.wavfile.write(
                file_name, rate=model.config.sampling_rate, data=output)
            logger.info("Text to Speech done")

        except Exception as e:
            logger.error(f"Error while TTS:{str(e)}")
            raise typer.Exit(1)
    
    def start(self):
        try:
            logger.info("Starting Pipeline")
            file = self.audio_name[:-4]
            # transcript = self.transcibe(self.input_path + self.audio_name)
            # translated_text = self.translate(transcript, self.language)
            
            # self.english_srt(transcript, self.input_path + self.audio_name)
            
            self.translated_sub(file, self.language)
            translated_transcript = " ".join(self.translated)
            
            self.tts(file, translated_transcript, languages[self.language]["tts"],languages[self.language][self.gender])
            self.translated.clear()
            logger.info("Pipeline Done")
        except Exception as e:
            logger.error(f"Error while running pipeline:{str(e)}")
            raise typer.Exit(1)


def seconds_to_subtitle_time(seconds):
        hours = int(seconds // 3600)
        seconds %= 3600
        minutes = int(seconds // 60)
        seconds %= 60
        milliseconds = int((seconds - int(seconds)) * 1000)
        seconds = int(seconds)

        return f"{hours:02d}:{minutes:02d}:{seconds:02d},{milliseconds:03d}"


def transcibe(input_path, audio_name):
    try:
        logger.info("Starting Transcription")
        try:
            model_size = 'large-v2'
            #use float16 for gpu
            model = WhisperModel(model_size, device='cpu', compute_type='int8')
        except Exception as e:
            logger.error(f'{e} thrown while loading whisper model')
            raise typer.Exit(1)
        segments, _ = model.transcribe(input_path+audio_name, beam_size=5)
        logger.info("Transciption Done")

        i=0
        srt=""
        output_file = f"{root_dir}/input/"+ audio_name[:-4] +".srt"
        for segment in segments:
            text = segment.text.lstrip()
            i=i+1
            srt = srt+"%i\n%s --> %s\n%s"%(i,seconds_to_subtitle_time(segment.start), seconds_to_subtitle_time(segment.end), text)+"\n\n"


        with open(output_file, 'x+', encoding='utf-8') as f:
            f.write(srt)
        return segments
        

    except Exception as e:
        logger.error(f"Error occured while transcribing text:{str(e)}")
        raise typer.Exit(1)
        


def get_gender(audioname,input_path):
    logger.info("Getting gender")
    device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    path = f"{input_path}{audioname}"
    pipe = pipeline(model="alefiury/wav2vec2-large-xlsr-53-gender-recognition-librispeech", trust_remote_code=True,device=device)
    result = pipe(path)
    logger.info(result[0]['label'])
    result = result[0]['label']
    return result
    
def process(input_path,audio_name,lang,gender):
    logger.info(f"{lang} process started")
    Pipeline(input_path,audio_name,lang,gender).start()
    logger.info(f"{lang} process completed")

def multi_process(input_path,audio,langs):
    start_time = time.time()
    logger.info("Multiprocessing started")
    
    gender = get_gender(audio,input_path)
    transcript = transcibe(input_path, audio)
    
    with Pool(processes=len(langs)) as pool:
        pool.starmap(process, zip(repeat(input_path),repeat(audio),langs,repeat(gender)))
    
    # for lang in langs:
    #     process(input_path,audio,lang)
    
    end_time = time.time()
    duration = end_time - start_time
    logger.info(f"Multiprocessing completed and time taken {duration}") 
    

    