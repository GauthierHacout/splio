## Splio test

The goal of this project is to be able to get real time informations about velib station around Splio's HQ. For this we use opendata API.
https://opendata.paris.fr/explore/dataset/velib-disponibilite-en-temps-reel/information/

**There is two version of this project :**

**1-** Every minute the informations of every stations in a 500m radius are outputed in the console

**2-** The informations of every stations in a 500m radius is displayed though a web page at
    __http://localhost:3000/__
    
To launch the project you can use ```go build main.go``` and then execute the compiled project by using a command-line argument :
``./main -output=console``

Ouput can be either ``console`` or ``web`` (it's web by default if you don't put anything)


### Potential improvements 
 - Right now we need to refresh the webpage to have an update on the information, we could implement AJAX request every minute to prevent having to do that and still getting the updated data for example. Or we could try implementing a websocket for example.
 - We could introduce the adress of the velib station in addition of the informations so people would know where to go if they are not familiar with Splio's HQ surronding (by using Google Maps APIs for example)
 - We could add the possibility for the user to choose the radius around Splio's HQ