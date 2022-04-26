# **Web crawler tool**

**1. Pre requirements**
Install docker on your local machine or the deployment environment

**2. How to run**

**Step 1:** Build the docker image by using the following command

```
docker build -t web-crawler .
```
**Step 2:** Run the service by using the following command

```
docker run -d -p 8000:8000 web-crawler
```

PS: you can select your own port mapping by modifying the port on `config.yml`

**Step 3: **Navigate to the following url on your favourite browser

```
localhost:8000\index
```

**3. Future improvements**
Improving the platform to integrate a sentiment analysis core so that it can grasp an idea about a particular website. Further, it can be used to get actual consumer reviews of a particular product or a service through social media and blog crawling.
Taking another step further, this can be improved with NLP tools to process integrated audio and video files to analyze information about a web page.