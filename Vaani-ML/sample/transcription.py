from faster_whisper import WhisperModel
# import logger

def seconds_to_subtitle_time(seconds):
        hours = int(seconds // 3600)
        seconds %= 3600
        minutes = int(seconds // 60)
        seconds %= 60
        milliseconds = int((seconds - int(seconds)) * 1000)
        seconds = int(seconds)

        return f"{hours:02d}:{minutes:02d}:{seconds:02d},{milliseconds:03d}"

model_size = 'large-v2'
    #use float16 for gpu
model = WhisperModel(model_size, device='cpu', compute_type='int8')



segments, _ = model.transcribe('E:\Vaani-ML\data\input\input.mp3', beam_size=5)
# logger.info("Transciption Done")

i=0
srt=""
output_file = f"E:\Vaani-ML\data\input\input"+".srt"
for segment in segments:
    text = segment.text.lstrip()
    i=i+1
    srt = srt+"%i\n%s --> %s\n%s"%(i,seconds_to_subtitle_time(segment.start), seconds_to_subtitle_time(segment.end), text)+"\n\n"


# with open(output_file, 'x+', encoding='utf-8') as f:
#     f.write(srt)

# print(seconds_to_subtitle_time(segment.start), seconds_to_subtitle_time(segment.end))