GHRSHR - Go Hot Reload and Scheduler High Rate 

for running cd mac and ./run.sh

for hot reaload ./reload

# building

docker build --tag mac .

docker run -p 4000:4000 mac

# inspect

docker run -ti -p 4000:4000 mac /bin/bash