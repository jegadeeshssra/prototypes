document.addEventListener("DOMContentLoaded", () => {
    const productsContainer = document.getElementById("products-container");
    const rawJsonContainer = document.getElementById("raw-json");
    const toggleJsonBtn = document.getElementById("toggle-json-btn");

    // Fetch products from the backend API
    fetch("http://localhost:8090/products/")
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            // Display raw JSON
            rawJsonContainer.textContent = JSON.stringify(data, null, 2);
            
            // Clear loading text
            productsContainer.innerHTML = "";

            // Check if data is an array
            if (Array.isArray(data) && data.length > 0) {
                // Iterate over the products and create HTML elements for each
                data.forEach(product => {
                    const productCard = document.createElement("div");
                    productCard.className = "product-card";

                    const title = document.createElement("div");
                    title.className = "product-title";
                    // Fallbacks in case the fields are named differently
                    title.textContent = product.Name || product.name || product.Title || product.title || `Product ID: ${product.ID || product.id}`;

                    const details = document.createElement("div");
                    details.className = "product-details";
                    details.textContent = product.Description || product.description || "No description available.";

                    const price = document.createElement("div");
                    price.className = "product-price";
                    price.textContent = `$${product.Price || product.price || "0.00"}`;

                    productCard.appendChild(title);
                    productCard.appendChild(details);
                    productCard.appendChild(price);

                    productsContainer.appendChild(productCard);
                });
            } else {
                productsContainer.innerHTML = "<div class='loading'>No products found.</div>";
            }
        })
        .catch(error => {
            console.error("Error fetching products:", error);
            productsContainer.innerHTML = `<div class="error">Failed to load products: ${error.message}</div>`;
        });

    // Toggle raw JSON view
    toggleJsonBtn.addEventListener("click", () => {
        if (rawJsonContainer.style.display === "block") {
            rawJsonContainer.style.display = "none";
            toggleJsonBtn.textContent = "Show Raw JSON";
        } else {
            rawJsonContainer.style.display = "block";
            toggleJsonBtn.textContent = "Hide Raw JSON";
        }
    });
});
