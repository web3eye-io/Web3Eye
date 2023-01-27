FROM coastlinesss/image-converter:0.0.1

COPY ./ /code
WORKDIR /code
RUN pip install  -r /code/requirements.txt

# Load the operator ahead of time
RUN python3 init.py
CMD ["python3","main.py"]