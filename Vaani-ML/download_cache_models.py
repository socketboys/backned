from transformers import AutoProcessor, SeamlessM4Tv2Model
from faster_whisper import WhisperModel
from transformers import AutoModelForSeq2SeqLM
from transformers import NllbTokenizerFast
from transformers.pipelines import pipeline
from loguru import logger

def main():
    try:
        # Your existing code here
        logger.info("Transcribing model")
        model_size = 'large-v2'
        transcript_model = WhisperModel(model_size, device='cpu', compute_type='int8')
        logger.info("Translate model")
        translate_tokenizer = NllbTokenizerFast.from_pretrained("facebook/nllb-200-distilled-600M")
        translate_model = AutoModelForSeq2SeqLM.from_pretrained("facebook/nllb-200-distilled-600M")
        logger.info("TTS model")
        processor = AutoProcessor.from_pretrained("facebook/seamless-m4t-v2-large")
        model = SeamlessM4Tv2Model.from_pretrained("facebook/seamless-m4t-v2-large")
        logger.info("Gender model")
        pipe = pipeline(model="alefiury/wav2vec2-large-xlsr-53-gender-recognition-librispeech", trust_remote_code=True)

        # Your further code/logic goes here

        logger.success("Script executed successfully.")

    except Exception as e:
        logger.error(f"An error occurred: {e}")
        # Optionally, you can raise the exception again if you want the script to terminate with an error
        # raise e

if __name__ == "__main__":
    main()
