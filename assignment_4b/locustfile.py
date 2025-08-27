from locust import FastHttpUser, task, between, events
import random
import json
import logging

# Set up logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Track metrics for GET vs POST ratio
class Metrics:
    def __init__(self):
        self.get_count = 0
        self.post_count = 0
        self.users_spawned = 0
    
    def add_get(self):
        self.get_count += 1
    
    def add_post(self):
        self.post_count += 1
    
    def add_user(self):
        self.users_spawned += 1
    
    def get_ratio(self):
        if self.post_count == 0:
            return "∞ (no POSTs yet)"
        return f"{self.get_count / self.post_count:.2f}:1"

metrics = Metrics()

# Sample product data for POST requests
def generate_product():
    product_id = random.randint(1, 999999)
    return {
        "product_id": product_id,
        "sku": f"SKU-{random.randint(100, 999)}-{random.choice(['A', 'B', 'C'])}{random.randint(10, 99)}",
        "manufacturer": random.choice(["TechCorp", "Innovate Inc.", "Quality Products", "NextGen", "PrimeManufacturing"]),
        "category_id": random.randint(1, 100),
        "weight": random.randint(100, 5000),
        "some_other_id": random.randint(1, 10000)
    }

# Locust User class - FastHttpUser for better performance
class ProductAPIUser(FastHttpUser):
    wait_time = between(1, 3)  # Wait between 1-3 seconds between tasks
    connection_timeout = 10.0
    network_timeout = 10.0
    
    def on_start(self):
        """Setup when each user starts"""
        logger.info("User started")
        metrics.add_user()
        logger.info(f"Users spawned so far: {metrics.users_spawned}")
        
        # We already know a product with ID 12345 exists
        self.known_product_id = 12345
        self.created_products = []
    
    @task(3)  # Weight of 3 for GET requests
    def get_product(self):
        """Task to GET a product - 3 times more frequent than POST"""
        # Choose between known product and one we've created
        if not self.created_products:
            product_id = self.known_product_id
        else:
            # 70% chance to use known product, 30% chance to use one we created
            if random.random() < 0.7 or not self.created_products:
                product_id = self.known_product_id
            else:
                product_id = random.choice(self.created_products)
                
        # Make the GET request
        with self.client.get(f"/products/{product_id}", catch_response=True) as response:
            if response.status_code == 200:
                response.success()
            elif response.status_code == 404:
                # Not found is valid for IDs we haven't created
                response.success()
            else:
                response.failure(f"Unexpected status code: {response.status_code}")
        
        metrics.add_get()
    
    @task(1)  # Weight of 1 for POST requests
    def create_product(self):
        """Task to POST a product - 1/3 as frequent as GET"""
        product = generate_product()
            
        with self.client.post("/products", json=product, catch_response=True) as response:
            if response.status_code == 201:
                # Store the created product ID for future GET requests
                self.created_products.append(product["product_id"])
                if len(self.created_products) > 10:  # Keep list manageable
                    self.created_products = self.created_products[-10:]
                response.success()
            else:
                response.failure(f"Unexpected status code: {response.status_code}")
        
        metrics.add_post()

# Event handlers with compatible signatures
@events.test_start.add_listener
def on_test_start(*args, **kwargs):
    logger.info("Test is starting")
    logger.info("Will simulate 50 users with 10 users/sec spawn rate")
    logger.info("Target GET:POST ratio is 3:1")
    
    # Reset metrics at test start
    metrics.get_count = 0
    metrics.post_count = 0
    metrics.users_spawned = 0

@events.test_stop.add_listener
def on_test_stop(*args, **kwargs):
    logger.info(f"Test finished. GET:POST ratio = {metrics.get_ratio()}")
    logger.info(f"GET requests: {metrics.get_count}, POST requests: {metrics.post_count}")
    logger.info(f"Total users spawned: {metrics.users_spawned}")