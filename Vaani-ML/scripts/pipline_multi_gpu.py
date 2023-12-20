import warnings
import sys
import os
from re import search
import typer
import scipy
import torch
from transformers import MBartForConditionalGeneration, MBart50TokenizerFast
from transformers import VitsModel, AutoTokenizer
from whisper import load_model, load_audio
from whisper.utils import get_writer
import ssl
from config import root_dir, languages,gpu_devices
from loguru import logger
from multiprocessing import Pool
import time
from itertools import repeat
# import ray
# from ray.util.multiprocessing import Pool
import torch.multiprocessing as mp

mp.set_start_method('spawn', force=True)

ssl._create_default_https_context = ssl._create_unverified_context

logger.add("./logs/{time}.log", level="TRACE", rotation="100 MB")

logger.info(f"Running on {sys.platform}")


class Pipeline:
    def __init__(self,input_path,audio_name,language):
        logger.info("Class initialized")
        self.input_path = input_path
        self.audio_name = audio_name
        self.language = language
        # self.device = torch.device('cuda' if torch.cuda.is_available() else 'cpu')
        self.device = torch.device('mps')
        logger.info(f"Device: {self.device}")
        self.translated = list()

    def transcibe(self,audio):
        try:
            logger.info("Starting Transcription")
            try:
                model = load_model("base")
            except Exception as e:
                logger.error(f'{e} thrown while loading whisper model')
                raise typer.Exit(1)
            input_audio = load_audio(audio)
            transcript = model.transcribe(input_audio, fp16=False)
            logger.info("Transciption Done")
            return transcript
        except Exception as e:
            logger.error(f"Error occured while transcribing text:{str(e)}")
            raise typer.Exit(1)
    
    def translate(self,text,lang):
        try:
            translate_model = MBartForConditionalGeneration.from_pretrained(
        "facebook/mbart-large-50-one-to-many-mmt").to(self.device)
            tokenizer = MBart50TokenizerFast.from_pretrained(
        "facebook/mbart-large-50-one-to-many-mmt", src_lang="en_XX")
        except Exception as e:  
            logger.error(f"Error occurred while loading the model/tokenizer: {str(e)}")
            raise typer.Exit(1)
        try:
            logger.info(f"Translating in {lang}")
            model_inputs = tokenizer(text, return_tensors="pt").to(self.device)

            generated_tokens = translate_model.generate(
                **model_inputs,
                forced_bos_token_id=tokenizer.lang_code_to_id[lang]
            )
            text = tokenizer.batch_decode(generated_tokens, skip_special_tokens=True)
            logger.info("Translation Done")
            return text[0]
        except Exception as e:
            logger.error(f"Error while translating text:{str(e)}")
            raise typer.Exit(1)
                    
    def english_srt(self,transcript, audio):
        # input_audio = load_audio(audio)
        try:
            srt_writer = get_writer("srt", f"{root_dir}/input/")
            # srt_writer = get_writer("srt", "/")
            srt_writer(transcript, audio)
        except Exception as e:
            logger.error(f"Error while writing the english subtitles:{str(e)}")
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
            
    def tts(self,file, text, lang):
        try:
            logger.info("Starting Text to Speech")
            file_name = f"{root_dir}/audio/" + file + "_" + lang[-3:-1] + ".wav"
            try:
                model = VitsModel.from_pretrained(lang).to(self.device)
                tts_tokenizer = AutoTokenizer.from_pretrained(lang)
            except Exception as e:
                logger.error(f"Error while loading TTS model:{str(e)}")
                raise typer.Exit(1)
            
            inputs = tts_tokenizer(text, return_tensors="pt").to(self.device)
            with torch.no_grad():
                output = model(**inputs).waveform
            scipy.io.wavfile.write(
                file_name, rate=model.config.sampling_rate, data=output.cpu().numpy().squeeze())
            logger.info("Text to Speech done")

        except Exception as e:
            logger.error(f"Error while TTS:{str(e)}")
            raise typer.Exit(1)
    
    def start(self):
        try:
            logger.info("Starting Pipeline")
            file = self.audio_name[:-4]
            transcript = self.transcibe(self.input_path + self.audio_name)
            # translated_text = self.translate(transcript, self.language)
            
            self.english_srt(transcript, self.input_path + self.audio_name)
            
            self.translated_sub(file, self.language)
            translated_transcript = " ".join(self.translated)
            
            self.tts(file, translated_transcript, languages[self.language]["tts"])
            self.translated.clear()
            logger.info("Pipeline Done")
        except Exception as e:
            logger.error(f"Error while running pipeline:{str(e)}")
            raise typer.Exit(1)
        
    
def process(input_path, audio_name, lang, gpu_device):
    torch.cuda.set_device(gpu_device)
    logger.info(f"{lang} process started on GPU {gpu_device}")
    Pipeline(input_path, audio_name, lang).start()
    logger.info(f"{lang} process completed on GPU {gpu_device}")


def multi_process(input_path,audio,langs):
    start_time = time.time()
    logger.info("Multiprocessing started")
    
    processes = []
    for i, lang in enumerate(langs):
        p = mp.Process(target=process, args=(input_path, audio, lang, gpu_devices[i % len(gpu_devices)]))
        processes.append(p)
        p.start()

    for p in processes:
        p.join()
         
    end_time = time.time()
    duration = end_time - start_time
    logger.info(f"Multiprocessing completed and time taken {duration}") 
    
    
    