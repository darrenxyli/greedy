#!/bin/bash
# Todo: add judgement of if the service has started

echo ""
echo "======================================================="
echo "||                                                   ||"
echo "||                        Greedy                     ||"
echo "||        ------------------------------------       ||"
echo "||        Distrubuted Crawl,  Mulitple Engines       ||"
echo "||        Mulitple Platforms,   Easily Monitor       ||"
echo "||        ------------------------------------       ||"
echo "||                                                   ||"
echo "||        WRITTEN FOR HUMAN   CODE BY DARRENXYLI     ||"
echo "||        github:darrenxyli  site:darrenxyli.com     ||"
echo "||                                                   ||"
echo "======================================================="
echo ""

# env params
if [ -z "${OUTPUT_PREFIX}" ]; then
    OUTPUT_PREFIX=/tmp
fi
mkdir -p ${OUTPUT_PREFIX}
echo -n "Set the log folder"
for i in 0 1 2 3 4; do
    echo -n "."
    sleep 0.1
done
echo " OK"

# start redis-server
OUTPUT_FILE=${OUTPUT_PREFIX}/redis-server-output-test.log
redis-server > ${OUTPUT_FILE} 2>&1 &
echo -n "start redis server"
for i in 0 1 2 3 4; do
    echo -n "."
    sleep 0.1
done
echo " OK"

# start celery workers
OUTPUT_FILE=${OUTPUT_PREFIX}/celery-output-test.log
LOG_FILE=${OUTPUT_PREFIX}/celery-log-test.log
#celeryd --loglevel=INFO > ${OUTPUT_FILE} 2>&1 &
celery multi start worker1 worker2 worker3 worker4 --loglevel=INFO > ${OUTPUT_FILE} 2>&1 &
echo -n "start celery workers"
for i in 0 1 2 3 4; do
    echo -n "."
    sleep 0.1
done
echo " OK"

# start mongoDB
OUTPUT_FILE=${OUTPUT_PREFIX}/mongoDB-output-test.log
mongod > ${OUTPUT_FILE} 2>&1 &
echo -n "start mongoDB"
for i in 0 1 2 3 4; do
    echo -n "."
    sleep 0.1
done
echo " OK"

# start celery-flower
OUTPUT_FILE=${OUTPUT_PREFIX}/celery-flower-output-test.log
celery flower --broker=redis://127.0.0.1:6379/0 > ${OUTPUT_FILE} 2>&1 &
echo -n "start flower monitoring"
for i in 0 1 2 3 4; do
    echo -n "."
    sleep 0.1
done
echo " OK"

# start scrapyd
OUTPUT_FILE=${OUTPUT_PREFIX}/scrapyd-output-test.log
(cd ~ && scrapyd > ${OUTPUT_FILE} 2>&1 &)
echo -n "start scrapyd service"
for i in 0 1 2 3 4; do
    echo -n "."
    sleep 0.1
done
echo " OK"

# deploy the scrapy
OUTPUT_FILE=${OUTPUT_PREFIX}/scrapy-output-test.log
(cd Sunflower/ && scrapyd-deploy Sunflower -p Sunflower > ${OUTPUT_FILE} 2>&1 &)
echo -n "deploy the scrapy project"
for i in 0 1 2 3 4; do
    echo -n "."
    sleep 0.1
done
echo " OK"

# start heimdallr monitor
OUTPUT_FILE=${OUTPUT_PREFIX}/heimdallr-output-test.log
(cd heimdallr && node app.js > ${OUTPUT_FILE} 2>&1 &)
echo -n "start Heimdallr System Monitor"
for i in 0 1 2 3 4; do
    echo -n "."
    sleep 0.1
done
echo " OK"
