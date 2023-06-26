import json
import os
import uuid
import urllib3
import logging
from confluent_kafka import Consumer, Producer
from pkg.model.resnet50 import Resnet50
from pkg.utils import imggetter
from pkg.utils import imgcheck
from pkg.utils import oss
from pkg.utils import config   
 # Parse the command line.
cConfig = {'bootstrap.servers': 'kafka-headless:9092',
        'group.id': 'python_default',
        'auto.offset.reset': 'earliest',
        'client.id': uuid.uuid4()}
pConfig = {'bootstrap.servers': 'kafka-headless:9092'}

# Create Consumer instance
consumer = Consumer(cConfig)
producer = Producer(pConfig)
pTopic = "image-converter-output"
cTopic = "image-converter-input"

def reset_offset(consumer, partitions):
    consumer.assign(partitions)

consumer.subscribe([cTopic], on_assign=reset_offset)

# Poll for new messages from Kafka and print them.
while True:
    msg = consumer.poll(60.0)
    print(msg)