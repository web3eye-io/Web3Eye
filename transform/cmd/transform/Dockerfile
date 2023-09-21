FROM coastlinesss/transform-base:0.0.1

COPY transform /usr/local/bin/transform
COPY model /pycode

WORKDIR /pycode
RUN pip install --no-cache-dir -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple

WORKDIR /usr/local/bin/
# CMD ./transform run
CMD sleep 10000000000000000