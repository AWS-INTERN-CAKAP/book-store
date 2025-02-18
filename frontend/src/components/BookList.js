import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';

function BookList() {
  const [books, setBooks] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchBooks();
  }, []);

  const fetchBooks = async () => {
    try {
      const response = await axios.get(`${process.env.REACT_APP_API_URL}/books`);
      setBooks(response.data);
      setError(null);
    } catch (error) {
      console.error('Error fetching books:', error);
      setError('Failed to fetch books. Please try again later.');
    }
  };

  const deleteBook = async (id) => {
    if (window.confirm('Are you sure you want to delete this book?')) {
      try {
        await axios.delete(`${process.env.REACT_APP_API_URL}/books/${id}`);
        fetchBooks();
      } catch (error) {
        console.error('Error deleting book:', error);
        setError('Failed to delete book. Please try again later.');
      }
    }
  };

  if (error) {
    return (
      <div className="text-center mt-8">
        <p className="text-red-500">{error}</p>
        <button
          onClick={fetchBooks}
          className="mt-4 bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
        >
          Retry
        </button>
      </div>
    );
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Book List</h1>
        <Link to="/add" className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">
          Add New Book
        </Link>
      </div>
      {books.length === 0 ? (
        <p className="text-center text-gray-500 mt-8">No books available.</p>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {books.map((book) => (
            <div key={book.id} className="bg-white rounded-lg shadow-md overflow-hidden">
              {book.path_gambar && (
                <img
                  src={`${process.env.REACT_APP_API_URL}/uploads/${book.path_gambar}`}
                  alt={book.namaBuku}
                  className="w-full h-48 object-cover"
                  onError={(e) => {
                    e.target.src = 'https://via.placeholder.com/300x200?text=No+Image';
                  }}
                />
              )}
              <div className="p-4">
                <h2 className="text-xl font-semibold mb-2">{book.namaBuku}</h2>
                <p className="text-gray-600 mb-2">Price: ${book.harga}</p>
                <p className="text-gray-600 mb-4">{book.deskripsi}</p>
                <div className="flex justify-between">
                  <Link
                    to={`/edit/${book.id}`}
                    className="bg-yellow-500 text-white px-3 py-1 rounded hover:bg-yellow-600"
                  >
                    Edit
                  </Link>
                  <button
                    onClick={() => deleteBook(book.id)}
                    className="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600"
                  >
                    Delete
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default BookList;