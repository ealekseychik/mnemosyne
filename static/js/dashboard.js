// File: static/js/dashboard.js

document.getElementById('logoutButton').addEventListener('click', () => {
    localStorage.removeItem('token');
    window.location.href = '/static/html/login.html';
});

async function fetchBooks() {
    const token = localStorage.getItem('token');
    const response = await fetch('/admin/books', {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`,
        },
    });

    const data = await response.json();

    if (response.ok) {
        displayBooks(data);
    } else {
        alert(data.error);
    }
}

function displayBooks(books) {
    const bookList = document.getElementById('bookList');
    bookList.innerHTML = '';

    books.forEach(book => {
        const bookItem = document.createElement('div');
        bookItem.className = 'book-item';
        bookItem.innerHTML = `
            <h3>${book.name}</h3>
            <p>Author: ${book.author}</p>
            <p>Borrowed By: ${book.borrowed_by || 'None'}</p>
            <button onclick="sendReminder('${book.guid}')">Send Reminder</button>
            <button onclick="editBook('${book.guid}')">Edit</button>
            <button onclick="deleteBook('${book.guid}')">Delete</button>
        `;
        bookList.appendChild(bookItem);
    });
}

async function sendReminder(guid) {
    const token = localStorage.getItem('token');
    const response = await fetch(`/admin/book/${guid}/reminder`, {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${token}`,
        },
    });

    const data = await response.json();

    if (response.ok) {
        alert('Reminder sent successfully');
    } else {
        alert(data.error);
    }
}

async function editBook(guid) {
    // Implement edit book functionality here
}

async function deleteBook(guid) {
    const token = localStorage.getItem('token');
    const response = await fetch(`/admin/book/${guid}`, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${token}`,
        },
    });

    if (response.ok) {
        fetchBooks();
    } else {
        const data = await response.json();
        alert(data.error);
    }
}

fetchBooks();
