import React from 'react';
import { Link } from 'react-router-dom';

function Navbar() {
  return (
    <nav className="bg-blue-600 text-white shadow-lg">
      <div className="container mx-auto px-4">
        <div className="flex justify-between items-center h-16">
          <Link to="/" className="text-2xl font-bold">
            Book Store
          </Link>
          <div className="space-x-4">
            <Link to="/" className="hover:text-blue-200">
              Home
            </Link>
            <Link to="/add" className="hover:text-blue-200">
              Add Book
            </Link>
          </div>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;