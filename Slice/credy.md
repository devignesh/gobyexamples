Q1. Output like this,
    start, First middleware
    End, First middleware

    start, Second middleware
    End, Second middleware

    start, Third middleware
    End, Third middleware

Q2. 
    even_squres = [i*i for i in range(100) if i % 2 == 0 ]

    o/p:
    devignesh@devignesh:~/go/src/gobyexamples$ python Slice/samp.py 
n = [0, 4, 16, 36, 64, 100, 144, 196, 256, 324, 400, 484, 576, 676, 784, 900, 1024, 1156, 1296, 1444, 1600, 1764,
 1936, 2116, 2304, 2500, 2704, 2916, 3136, 3364, 3600, 3844, 4096, 4356, 4624, 4900, 5184, 5476, 5776, 6084, 6400
, 6724, 7056, 7396, 7744, 8100, 8464, 8836, 9216, 9604]

Q3: What should be the HTTP response code for the following situations:
    a) If a request is sent on a URL which does not exist on the application. ----- 404 Not found
    b) If a request is sent on a URL to create an object on the application, and it is successfully created. ----- 201 Created 
    c) If a request is sent on a URL to get details of an object and the detail is sent successfully. -----  200 Ok
    d) If a form is submitted with some invalid input. --- 406 Not acceptable

Q4: Design a movie and actor database, where you can store movies and actor names, and which actor appeared in which movie. Write the models you will use and foreign key relationships.

from django.db import models

# Actor model
class Actor(models.Model):
    
    actor_name = models.TextField(max_length=285, blank=True, null=True)
    published_date = models.DateTimeField(auto_now_add=True, null=True, blank=True)
    updated_at = models.DateTimeField(null=True, blank=True)

    def __str__(self):
        return self.actor_name
    
#movie model
class Movie(models.Model):
    
    movie_name = models.CharField(max_length=285, blank=True, null=True)
    actor = models.ForeignKey(Actor, null=True, blank=True, on_delete=models.SET_NULL)
    director = models.CharField(max_length=285, blank=True, null=True)
    language = models.CharField(max_length=285, blank=True, null=True)
    relesed_time = models.DateTimeField(null=True, blank=True)
    created_at = models.DateTimeField(auto_now_add=True,null=True, blank=True)
    
    def __str__(self):
        return self.movie_name