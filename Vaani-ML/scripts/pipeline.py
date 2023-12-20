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
from config import root_dir
from loguru import logger

ssl._create_default_https_context = ssl._create_unverified_context

logger.add("../logs/{time}.log", level="TRACE", rotation="100 MB")

logger.info(f"Running on {sys.platform}")

device = torch.device('cuda' if torch.cuda.is_available() else 'cpu')

try:
    translate_model = MBartForConditionalGeneration.from_pretrained(
        "facebook/mbart-large-50-one-to-many-mmt").to(device)
    tokenizer = MBart50TokenizerFast.from_pretrained(
        "facebook/mbart-large-50-one-to-many-mmt", src_lang="en_XX")
except Exception as e:
    logger.error(f"Error occurred while loading the model/tokenizer: {str(e)}")
    raise typer.Exit(1)

translated = list()

languages = {
    "hindi": {
        "translate": "hi_IN",
        "tts": "facebook/mms-tts-hin"
    },
    "marathi": {
        "translate": "mr_IN",
        "tts": "facebook/mms-tts-mar"
    },
    "bengali": {
        "translate": "bn_IN",
        "tts": "facebook/mms-tts-ben"
    },
    "tamil": {
        "translate": "ta_IN",
        "tts": "facebook/mms-tts-tam"
    },
    "telegu": {
        "translate": "te_IN",
        "tts": "facebook/mms-tts-tel"
    },
}

def check_device():
    if device == 'cpu':
        warnings.warn("You are using CPU.")
        logger.warning("You are using CPU")


def transcibe(audio):
    try:
        logger.info("Starting Transcription")
        try:
            model = load_model("base")
        except Exception as e:
            logger.error(f'{e} thrown while loading whisper model ')
            raise typer.Exit(1)
        input_audio = load_audio(audio)
        transcript = model.transcribe(input_audio, fp16=False)
        logger.info("Transciption Done")
        return transcript
    except Exception as e:
        logger.error(f"Error occured while transcribing text:{str(e)}")
        raise typer.Exit(1)


def translate(text, lang):
    try:
        # logger.info("Translating...")
        model_inputs = tokenizer(text, return_tensors="pt").to(device)

        generated_tokens = translate_model.generate(
            **model_inputs,
            forced_bos_token_id=tokenizer.lang_code_to_id[lang]
        )
        text = tokenizer.batch_decode(generated_tokens, skip_special_tokens=True)
        # logger.info("Translation Done")
        return text[0]
    except Exception as e:
        logger.error(f"Error while translating text:{str(e)}")
        raise typer.Exit(1)

def tts(file, text, lang):
    try:
        logger.info("Starting Text to Speech")
        file_name = f"{root_dir}/audio/" + file + "_" + lang[-3:-1] + ".wav"
        try:
            model = VitsModel.from_pretrained(lang).to(device)
            tts_tokenizer = AutoTokenizer.from_pretrained(lang)
        except Exception as e:
            logger.error(f"Error while loading TTS model:{str(e)}")
            raise typer.Exit(1)
        
        inputs = tts_tokenizer(text, return_tensors="pt").to(device)
        with torch.no_grad():
            output = model(**inputs).waveform
        scipy.io.wavfile.write(
            file_name, rate=model.config.sampling_rate, data=output.cpu().numpy().squeeze())
        logger.info("Text to Speech done")

    except Exception as e:
        logger.error(f"Error while TTS:{str(e)}")
        raise typer.Exit(1)

def english_srt(transcript, audio):
    # input_audio = load_audio(audio)
    try:
        srt_writer = get_writer("srt", f"{root_dir}/input/")
        # srt_writer = get_writer("srt", "/")
        srt_writer(transcript, audio)
    except Exception as e:
        logger.error(f"Error while writing the english subtitles:{str(e)}")
        raise typer.Exit(1)


def translated_sub(file, lang):
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

                    translated_text = translate(line, languages[lang]["translate"])
                    translated.append(translated_text)
                    outfile.write(translated_text + "\n\n")
        logger.info("Translated.")
        logger.info("Subtitles are created")
    except Exception as e:
        logger.error(f"Error while translating subtitles:{str(e)}")
        raise typer.Exit(1)


def pipeline(path, audio, langs):
    logger.info("Starting the pipeline")
    check_device()
    file = audio[:-4]
    transcript = transcibe(path + audio)
    # text = transcript["text"]
    english_srt(transcript, path + audio)
    for lang in langs:
        translated_sub(file, lang)
        translated_transcript = " ".join(translated)
        # translated_text = translate(text, languages[lang]["translate"])
        tts(file, translated_transcript, languages[lang]["tts"])
        translated.clear()
    logger.info("Pipeline finished")


# pipeline("converted.mp3.mp3", ["hindi"])