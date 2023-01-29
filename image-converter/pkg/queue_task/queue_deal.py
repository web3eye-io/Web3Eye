#!/usr/bin/env python

import json
import os
import uuid
from confluent_kafka import Consumer, Producer
from pkg.model.resnet50 import Resnet50
from pkg.utils import imggetter
from pkg.utils import imgcheck


def UUIDCheck(id) -> bool:
    try:
        uuid.UUID(id)
        return True
    except ValueError:
        return False


class VectorInfo(object):
    def __init__(self, id="", url="", vector=[], msg="", success=False) -> None:
        self.id = id
        self.url = url
        self.vector = vector
        self.msg = msg
        self.success = success


class VectorInfoEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, VectorInfo):
            return {
                'id': o.id,
                'url': o.url,
                'vector': o.vector,
                'msg': o.msg,
                'success': o.success,
            }
        return json.JSONEncoder.default(o)


def QueueDealImageURL2Vector():
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
    try:
        while True:
            msg = consumer.poll(60.0)
            vectorInfo = VectorInfo(success=False)

            # when have no data,producer flush
            if msg is None:
                continue
            elif msg.error():
                print("ERROR: %s".format(msg.error()))
            else:
                print(msg.value().decode('utf-8'), msg.key().decode('utf-8'))
                vectorInfo.url = msg.value().decode('utf-8')
                vectorInfo.id = msg.key().decode('utf-8')

                image_path, ok = imggetter.DownloadUrlImg(vectorInfo.url)
                if not ok:
                    vectorInfo.msg = "url format cannot parse,url: " + vectorInfo.url
                    producer.produce(pTopic, json.dumps(
                        vectorInfo, cls=VectorInfoEncoder), vectorInfo.id)
                    continue

                imgcheck.CheckImg(path=image_path)
                vectorInfo.vector = Resnet50().resnet50_extract_feat(img_path=image_path)
                os.remove(path=image_path)

                vectorInfo.success = True
                producer.produce(pTopic, json.dumps(
                    vectorInfo, cls=VectorInfoEncoder), vectorInfo.id)

    except KeyboardInterrupt:
        pass
    finally:
        # Leave group and commit final offsets
        consumer.close()
        producer.poll(10000)
        producer.flush()
