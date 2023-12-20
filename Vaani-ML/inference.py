import argparse
from typing import List
from scripts import pipeline_class,pipeline_ai
from config import root_dir

input_dir = f"{root_dir}/input/"

seamless = ["hindi","bengali","telugu"]

def run(langs:List, audioname:str):

    if any(lang in seamless for lang in langs):
        try:
            pipeline_class.multi_process(input_dir, audioname, langs)
        except Exception as e:
            print(f'{e} thrown from pipeline')
            exit(1)
    else:
        try:
            pipeline_ai.multi_process(input_dir, audioname, langs)
        except Exception as e:
            print(f'{e} thrown from pipeline')
            exit(1)

def main():
    parser = argparse.ArgumentParser(description="Process audio input with specified languages")
    parser.add_argument("--lang", nargs="+", required=True, help="List of languages")
    parser.add_argument("--audioname",required=True, help="Name of the audio file")

    args = parser.parse_args()

    if len(args.lang) == 0:
        parser.error("No languages detected")
    
    if any(lang in seamless for lang in args.lang):
        try:
            pipeline_class.multi_process(input_dir, args.audioname, args.lang)
        except Exception as e:
            print(f'{e} thrown from pipeline')
            exit(1)
    else:
        try:
            pipeline_ai.multi_process(input_dir, args.audioname, args.lang)
        except Exception as e:
            print(f'{e} thrown from pipeline')
            exit(1)

if __name__ == "__main__":
    main()