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
from transformers.pipelines import pipeline
import ssl
from config import root_dir, languages,gender
from loguru import logger
from multiprocessing import Pool
import time
from itertools import repeat
import torchaudio
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
            
    def tts(self,file, lang,gender):
        path = f"{self.input_path}{self.audio_name}"
        try:
            logger.info("Starting Speech to Speech")
            file_name = f"{root_dir}/audio/" + file + "_" + lang[-3:-1] + ".wav"
            try:
               processor = AutoProcessor.from_pretrained("facebook/seamless-m4t-v2-large")
               model = SeamlessM4Tv2Model.from_pretrained("facebook/seamless-m4t-v2-large").to(self.device)
            except Exception as e:
                logger.error(f"Error while loading TTS model:{str(e)}")
                raise typer.Exit(1)
            audio, orig_freq =  torchaudio.load(path)
            audio =  torchaudio.functional.resample(audio, orig_freq=orig_freq, new_freq=16_000) # must be a 16 kHz waveform array
            audio_inputs = processor(audios=audio, return_tensors="pt").to(self.device)
            output= model.generate(**audio_inputs, tgt_lang=lang,speaker_id=gender)[0].cpu().numpy().squeeze()
            scipy.io.wavfile.write(
                file_name, rate=14000, data=output)
            logger.info("Speech to Speech done")

        except Exception as e:
            logger.error(f"Error while TTS:{str(e)}")
            raise typer.Exit(1)
    
    def start(self):
        try:
            logger.info("Starting Pipeline")
            file = self.audio_name[:-4]
            self.tts(file, languages[self.language]["tts"],gender[self.gender])
            logger.info("Pipeline Done")
        except Exception as e:
            logger.error(f"Error while running pipeline:{str(e)}")
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
    
    
    with Pool(processes=len(langs)) as pool:
        pool.starmap(process, zip(repeat(input_path),repeat(audio),langs,repeat(gender)))
    
    # for lang in langs:
    #     process(input_path,audio,lang)
    
    end_time = time.time()
    duration = end_time - start_time
    logger.info(f"Multiprocessing completed and time taken {duration}") 
    

    