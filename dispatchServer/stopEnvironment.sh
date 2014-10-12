ps auxww | grep 'celery' | awk '{print $2}' | xargs kill -9
ps auxww | grep 'app.js' | awk '{print $2}' | xargs kill -9
ps auxww | grep 'redis' | awk '{print $2}' | xargs kill -9
ps auxww | grep 'mongo' | awk '{print $2}' | xargs kill -9
ps auxww | grep 'scrapyd' | awk '{print $2}' | xargs kill -9
rm -rf *.pyc
rm -rf ./Sunflower/build/
rm -rf dbs/
rm -rf eggs/
rm -rf logs/
rm -rf items/
rm -rf dump.rdb
rm -rf twistd.pid
rm -rf *.log
rm -rf *.pid
