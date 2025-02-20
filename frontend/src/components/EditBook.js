import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import axios from 'axios';
import CreatableSelect from "react-select/creatable";

function EditBook() {
  const navigate = useNavigate();
  const { id } = useParams();
  const [book, setBook] = useState({
    title: '',
    price: '',
    description: '',
    categories: [],
  });
  const [categories, setCategories] = useState([]);
  const [image, setImage] = useState(null);

  useEffect(() => {
    const fetchBook = async () => {
      try {
        const response = await axios.get(`${process.env.REACT_APP_API_URL}/books/${id}`);
        const data = response.data.data

        setBook({ 
          title: data.title,  
          price: data.price,
          description: data.description,
          categories: data.categories,
        });

      } catch (error) {
        console.error('Error fetching book:', error);
      }
    };

    const fetchCategories = async () => {
      try {
        const response = await axios.get(`${process.env.REACT_APP_API_URL}/categories`);
        const formattedCategories = response.data.data.map(cat => ({
          value: cat.id,
          label: cat.name,
        }));
        setCategories(formattedCategories);
      } catch (error) {
        console.error('Error fetching categories:', error);
      }
    };

    fetchBook();
    fetchCategories();
  }, []);

  const handleCategoryChange = (selectedOptions) => {
    const filteredCategories = categories.filter(category =>
      selectedOptions.some(option => option.value == category.value)
    );

    const formattedCategories = filteredCategories.map(category =>( {
      id: category.value,
      name: category.label
    }))
  

    setBook(prevBook => ({ ...prevBook, categories: formattedCategories }));
  };

  const handleCreateCategory = async (inputValue) => {
    try {
      
      const response = await axios.post(`${process.env.REACT_APP_API_URL}/categories`, {
        name: inputValue
      });

      const data = response.data.data
      const newCategory = { value: data.id, label: data.name }
      const bookCategory = {id: data.id, name: data.name}


      setCategories(prevCategories => [...prevCategories,newCategory]);
      setBook(prevBook => ({ ...prevBook, categories: [...prevBook.categories, bookCategory] }));

    } catch (error) {
      console.error('Error create category:', error)
    }
    
  };

  const handleChange = (e) => {
    setBook({ ...book, [e.target.name]: e.target.value });
  };

  const handleImageChange = (e) => {
    setImage(e.target.files[0]);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    console.log()
    const formData = new FormData();
    formData.append('title', book.title);
    formData.append('price', book.price);
    formData.append('description', book.description);
    formData.append('categories', book.categories.map((categories) => categories.id));
    if (image) {
      formData.append('image', image);
    }

    try {
      await axios.put(`${process.env.REACT_APP_API_URL}/books/${id}`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      navigate('/');
    } catch (error) {
      console.error('Error updating book:', error);
    }
  };

  return (
    <div className="max-w-2xl mx-auto">
      <h1 className="text-3xl font-bold mb-6">Edit Book</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700">Title</label>
          <input
            type="text"
            name="title"
            value={book.title}
            onChange={handleChange}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700">Price</label>
          <input
            type="number"
            name="price"
            value={book.price}
            onChange={handleChange}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700">Description</label>
          <textarea
            name="description"
            value={book.description}
            onChange={handleChange}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            rows="3"
            
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700">Categories</label>
          <CreatableSelect
            isMulti
            options={categories}
            onChange={handleCategoryChange}
            onCreateOption={handleCreateCategory}
            value={book.categories.map(cat => ({ value: cat.id, label: cat.name }))}
            className="mt-1"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700">Image</label>
          <input
            type="file"
            onChange={handleImageChange}
            className="mt-1 block w-full"
            accept="image/*"
          />
        </div>
        <div className="flex justify-end space-x-4">
          <button
            type="button"
            onClick={() => navigate('/')}
            className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600"
          >
            Cancel
          </button>
          <button
            type="submit"
            className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
          >
            Update Book
          </button>
        </div>
      </form>
    </div>
  );
}

export default EditBook;
