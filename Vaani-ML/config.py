root_dir = "../external" #prod
# root_dir = "./data"
languages = {
            "hindi": {
                "translate": "hin_Deva",
                "tts": "hin",
                "female": 3,
                "male": 1 
            },
        
            "bengali": {
                "translate": "ben_Beng",
                "tts": "ben",
                "female": 3,
                "male": 2  
            },
          
            "telugu": {
                "translate": "tel_Telu",
                "tts": "tel",
                "female": 3,
                "male": 2  
            },
             "assamese": {
                "translate": "asm_Beng",
                "tts": "asm",
            },
            "bodo": {
                "translate": "bod_Tibt",
                "tts": "bod",
            },
            "gujrati": {
                "translate": "guj_Gujr",
                "tts": "guj",
            },
             "kannada": {
                "translate": "kan_Knda",
                "tts": "kan",
            },
             "malyalam": {
                "translate": "mal_Mlym",
                "tts": "mal",
            },
             "marathi": {
                "translate": "mar_Deva",
                "tts": "mar",
            },
              "manipuri": {
                "translate": "mni_Beng",
                "tts": "man",
            },
               "odiya": {
                "translate": "ory_Orya",
                "tts": "odi",
            },
              "punjabi": {
                "translate": "pan_Guru",
                "tts": "pun",
            },
               "tamil": {
                "translate": "tam_Taml",
                "tts": "",
            },
           
           
        }
gender = {
            "female": 3,
            "male": 2   
        }
AIBharat_TTS = {
    'assamese':'as', 
    'bodo':'brx', 
    # 'hinglish':'en+hi', 
    'gujrati':'gu', 
    # 'hindi':'hi', 
    'kannada':'kn', 
    'malyalam':'ml', 
    'manipuri':'mni', 
    'marathi':'mr', 
    'odiya':'or', 
    'punjabi':'pa',
    # 'rajasthani':'raj', 
    'tamil':'ta', 
    # 'telugu':'te'
}
gpu_devices = [0, 1, 2] 