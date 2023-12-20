import wget
import os
import zipfile

def download_and_extract_model(model, download_link, output_dir):
    zip_file_path = os.path.join(output_dir, f'{model}.zip')
    wget.download(download_link + f'{model}.zip', out=output_dir)
    
    with zipfile.ZipFile(zip_file_path, 'r') as zip_ref:
        zip_ref.extractall(output_dir)

    os.remove(zip_file_path)

def main():
    link = 'https://github.com/AI4Bharat/Indic-TTS/releases/download/v1-checkpoints-release/'
    # models = ['as']
    models = ['as', 'brx', 'gu', 'kn', 'ml', 'mni', 'mr', 'or', 'pa', 'ta']

    output_directory = 'models/v1/'

    for model in models:
        download_and_extract_model(model, link, output_directory)

if __name__ == "__main__":
    main()

