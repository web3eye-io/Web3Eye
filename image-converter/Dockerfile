FROM python:3.9

COPY ./ /code
WORKDIR /code
RUN /usr/local/bin/python3 -m pip install --upgrade pip
RUN curl -s https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh | bash
RUN apt install git-lfs libgl1 python3-brlapi -y
RUN git lfs install
RUN pip install  -r /code/requirements.txt

# Load the operator ahead of time
RUN python3 init.py
CMD ["python3","cmd/image-converter/main.py"]