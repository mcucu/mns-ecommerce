// This file contains the JavaScript code for the web application, including functions to handle item selection and the "Place order" button click event.

document.addEventListener('DOMContentLoaded', function() {
    const itemList = document.getElementById('item-list');
    const placeOrderButton = document.getElementById('place-order');

    placeOrderButton.addEventListener('click', function() {
        const selectedItems = [];
        const checkboxes = itemList.querySelectorAll('input[type="checkbox"]:checked');

        checkboxes.forEach(function(checkbox) {
            selectedItems.push({
                id: checkbox.value,
                name: checkbox.dataset.name,
                price: parseFloat(checkbox.dataset.price),
                weight: parseFloat(checkbox.dataset.weight)
            });
        });

        if (selectedItems.length === 0) {
            alert('Please select at least one item to place an order.');
            return;
        }

        const orderData = {
            items: selectedItems,
            totalPrice: selectedItems.reduce((total, item) => total + item.price, 0),
            totalWeight: selectedItems.reduce((total, item) => total + item.weight, 0)
        };

        // Send order data to the server
        fetch('/api/orders', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(orderData)
        })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Order placed successfully!');
                // Optionally, reset the form or redirect
            } else {
                alert('Error placing order: ' + data.message);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('An error occurred while placing the order.');
        });
    });
});