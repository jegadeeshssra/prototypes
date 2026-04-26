// Simple script to fetch products from http://localhost:8090/products/ and display the JSON

document.addEventListener('DOMContentLoaded', function() {
    fetch('http://localhost:8090/products/')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            // Display the JSON data
            const productsDiv = document.getElementById('products');
            productsDiv.innerHTML = '<pre>' + JSON.stringify(data, null, 2) + '</pre>';
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
            const productsDiv = document.getElementById('products');
            productsDiv.innerHTML = 'Error: ' + error.message;
        });
});