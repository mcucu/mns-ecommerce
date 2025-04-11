// This file contains the JavaScript code for the web application, including functions to handle item selection and the "Place order" button click event.
const showSpinner = () => {
    $('#spinner').show();
}

const hideSpinner = () => {
    $('#spinner').hide();
}

const fetchProducts = async () => {
    try {
        const response = await $.ajax({
            url: apiUrl + '/products',
            method: 'GET',
            contentType: 'application/json'
        });

        const {
            success,
            message,
            data
        } = response;

        // Check if the response is successful
        if (!success) {
            alert(message);
            return;
        }

        // Check if data is empty
        if (data.length === 0) {
            alert('No products available.');
            return;
        }

        // Check if data is an array
        if (!Array.isArray(data)) {
            throw new Error('Invalid data format');
        }

        // Check if data contains the expected properties
        data.forEach(item => {
            if (!item.name || !item.price || !item.weight) {
                throw new Error('Invalid product data');
            }
        });

        // Return the fetched data
        console.log('Fetched products:', data);
        return data;
    } catch (error) {
        console.error('Error fetching products:', error);
        throw error;
    }
}

const postOrder = async (payload) => {
    try {
        const response = await $.ajax({
            url: apiUrl + '/orders',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(payload)
        });

        const {
            success,
            message,
            data
        } = response;

        // Check if the response is successful
        if (!success) {
            alert(message);
            return;
        }

        // Check if data is empty
        if (data.length === 0) {
            alert('No packages were created.');
            return;
        }

        console.log('Order created:', data);
        return data;
    } catch (error) {
        console.error('Error creating order:', error);
        throw error;
    }
}

const loadItems = async () => {
    showSpinner();
    let items = []
    try {
        items = await fetchProducts();
        if (!items) {
            alert('Failed to load products. Please try again.');
            hideSpinner();
            return;
        }
        if (items.length === 0) {
            alert('No products available.');
            hideSpinner();
            return;
        }
    } catch (error) {
        alert('Failed to load products. Please try again.');
        hideSpinner();
        return;
    }

    const itemList = $('#item-list .list-group');
    items.forEach(item => {
        itemList.append(`
            <li class="list-group-item">
                <input type="checkbox" class="item" data-name="${item.name}" data-price="${item.price}" data-weight="${item.weight}">
                ${item.name} - $${item.price} - ${item.weight}g
            </li>
        `);
    });

    hideSpinner();
}

const placeOrder = async () => {
    showSpinner();
    const selectedItems = [];
    $('.item:checked').each(function () {
        selectedItems.push({
            name: $(this).data('name'),
            price: $(this).data('price'),
            weight: $(this).data('weight')
        });
    });

    const payload = {
        items: selectedItems,
        total_weight: selectedItems.reduce((total, item) => total + item.weight, 0),
        total_price: selectedItems.reduce((total, item) => total + item.price, 0),
    }
    if (selectedItems.length === 0) {
        alert('Please select at least one item.');
        return;
    }

    let orderPackages = [];
    try {
        orderPackages = await postOrder(payload);
        console.log('Order packages:', orderPackages);
    } catch (error) {
        console.error('Error creating order:', error);
        hideSpinner();
        return;
    }

    if (orderPackages.length === 0) {
        alert('No packages were created.');
        return;
    }

    let resultHtml = '<h2>This order has the following packages:</h2>';
    orderPackages.forEach((pkg, index) => {
        resultHtml += `<h3>Package ${index + 1}</h3>`;
        resultHtml += `<p>Items: ${pkg.items.map(item => item.name).join(', ')}</p>`;
        resultHtml += `<p>Total Weight: ${pkg.total_weight}g</p>`;
        resultHtml += `<p>Total Price: $${pkg.total_price.toFixed(2)}</p>`;
        resultHtml += `<p>Courier Cost: $${pkg.courier_cost.toFixed(2)}</p>`;
    });
    $('#result').html(resultHtml);
    $('#clear-order').removeClass('disabled');
    hideSpinner();
}

const clearOrder = (obj) => {
    if ($(obj).hasClass('disabled')) {
        return;
    }
    $('.item').prop('checked', false);
    $('#result').empty();
    $(obj).addClass('disabled');
}

$(document).ready(function () {
    loadItems();

    $('#place-order').click(function () {
        placeOrder();
    });

    $('#clear-order').click(function () {
        clearOrder(this);
    });
});
