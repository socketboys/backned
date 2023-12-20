import pika
import json
from time import sleep
from scripts.pipeline_ai import multi_process
from inference import run


class Rabbit:
    producer_channel: pika.adapters.blocking_connection.BlockingChannel = None
    POOL_EXCHANGE = 'pool'
    POOL_ROUTING_KEY = 'poolkey'

    # space: Spaces

    def get_audio_type(self, language):
        if language == "hindi":
            return "_hi"
        elif language == "telugu":
            return "_tel"
        elif language == "bengali":
            return "_be"
        elif language == "assamese":
            return "_asm"
        elif language == "bodo":
            return "_bod"
        elif language == "gujrati":
            return "_guj"
        elif language == "kannada":
            return "_kan"
        elif language == "malyalam":
            return "_mal"
        elif language == "marathi":
            return "_mar"
        elif language == "manipuri":
            return "_mni"
        elif language == "odiya":
            return "_odi"
        elif language == "punjabi":
            return "_pan"
        elif language == "tamil":
            return "_tam"


    def callback(self, channel, method, properties, body):
        msg = json.loads(body)
        print(msg)
        run(msg["language"],msg["euid"])
        # sleep(120) # for testing

        channel.basic_ack(delivery_tag=method.delivery_tag)
        self.producer_channel.basic_publish(exchange=Rabbit.POOL_EXCHANGE, routing_key=Rabbit.POOL_ROUTING_KEY, body=body)

    def __init__(self):
        RABBIT_URL = 'amqp://guest:guest@localhost:5672'

        TRANSLATION_ROUTING_KEY = 'translationkey'
        TRANSLATION_QUEUE_NAME = 'translation_pipeline'
        TRANSLATION_EXCHANGE = 'translation'

        POOL_QUEUE_NAME = 'update_pool'

        # self.space = Spaces()

        self.connection = pika.BlockingConnection(pika.URLParameters(RABBIT_URL))

        self.producer_channel = self.connection.channel()
        self.producer_channel.queue_declare(queue=POOL_QUEUE_NAME, auto_delete=False, durable=True, arguments={'x-queue-type': 'quorum'})

        self.consumer_channel = self.connection.channel()
        self.consumer_channel.queue_declare(queue=TRANSLATION_QUEUE_NAME, auto_delete=False, durable=True, arguments={'x-single-active-consumer': True})
        self.consumer_channel.queue_bind(queue=TRANSLATION_QUEUE_NAME, exchange=TRANSLATION_EXCHANGE, routing_key=TRANSLATION_ROUTING_KEY)
        self.consumer_channel.basic_qos(prefetch_count=1)
        self.consumer_channel.basic_consume(TRANSLATION_QUEUE_NAME, on_message_callback=self.callback)

        print("Vaaani-ML starting to consume")

        self.consumer_channel.start_consuming()


if __name__ == "__main__":
    Rabbit()
