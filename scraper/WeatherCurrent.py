# SCRIPT FOR SCRAPING CURRENT WEATHER DATA
# DATA IS SCRAPED EVERY 20 MINS AND INSERTED IN MONGODDB - CURRENTWEATHER COLLECTION

# importing libraries
import requests
import json
import time
import os
import configparser

# error checking when connecting to MongoClient
try:
    from pymongo import MongoClient
except ImportError:
    raise ImportError('PyMongo is not installed')

# load config file
print('reading configurations')
config = configparser.ConfigParser()
config.read('config/scrapercfg.ini')
connectionsconfig = config['scraper']

def weather_current_main():
    urlWeatherCurrent = connectionsconfig['urlWeatherCurrent']
    urlWeatherCurrent = urlWeatherCurrent + "?lat=%s&lon=%s&appid=%s&units=metric"
    urlWeatherCurrent = urlWeatherCurrent % (
            connectionsconfig['lat'],
            connectionsconfig['lon'],
            connectionsconfig['api_key_current']) 
    response = requests.get(urlWeatherCurrent)

    response = requests.get(urlWeatherCurrent)
    data = response.text

    # testing to ensure the data was scraped
    if response.status_code != 200:
        print('Failed to get data:', response.status_code)

    # parsing response text to json format
    print('[*] Parsing response text')
    data = json.loads(response.text)

    # print("pushing data to mongodb with the functions")

    # inserting data to mongodb database
    print('[*] Pushing data to MongoDB ')
    cluster = MongoClient(connectionsconfig['uri'])
    db = cluster["Weather"]
    collection = db["CurrentWeather"]

    # inserting data in mongodb
    try:
        #creating index - datetime is unique 
        collection.create_index([('dt', -1)], unique=True)

        # inserting data 
        collection.insert_one(data)
        
    except Exception as ex:
        print(ex)
    else:
        print("Data inserted successfully")

    # close the connection
    finally:
        cluster.close()

    # weather info will be scraped every 20 minutes
    time.sleep(20 * 60)


while True:
    weather_current_main()
