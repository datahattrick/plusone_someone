import React, { useState, useEffect } from 'react';

interface UserSuggestion {
  id: number;
  name: string;
}

const MyForm: React.FC = () => {
  const [user, setUser] = useState('');
  const [message, setMessage] = useState('');
  const [suggestions, setSuggestions] = useState<UserSuggestion[]>([]);

  useEffect(() => {
    const fetchSuggestions = async () => {
      if (user.trim() === '') {
        setSuggestions([]);
        return;
      }

      try {
        const response = await fetch(`/api/v1/user/${user}`);
        if (response.ok) {
          const data = await response.json();
          setSuggestions(data);
        } else {
          console.error('Error fetching user suggestions');
        }
      } catch (error) {
        console.error('Error:', error);
      }
    };

    fetchSuggestions();
  }, [user]);

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    // Create an object with the user and message data
    const formData = { user, message };

    try {
      // Make a POST request to the app endpoint
      const response = await fetch('/api/v1/posts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      // Check the response status
      if (response.ok) {
        console.log('Data successfully submitted');
      } else {
        console.error('Error submitting data');
      }
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (


    <form onSubmit={handleSubmit}>
      <div>
        <input
          className="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none"
          type="text"
          id="user"
          value={user}
          onChange={(event) => setUser(event.target.value)}
        />
        {suggestions.length > 0 && (
          <ul>
            {suggestions.map((suggestion) => (
              <li key={suggestion.id} onClick={() => setUser(suggestion.name)}>
                {suggestion.name}
              </li>
            ))}
          </ul>
        )}
      </div>
      <div  className="mb-6">
        <textarea
         className="block w-full p-4 border border-gray-300 rounded-lg bg-gray-50 sm:text-md focus:ring-orange-200 focus:border-blue-500 "
          id="message"
          value={message}
          onChange={(event) => setMessage(event.target.value)}
        />
      </div>
      <button className="rounded-md bg-orange-400 px-3.5 py-2.5 font-semibold text-white shadow-sm hover:bg-orange-300 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 text-2xl" type="submit">Submit</button>
    </form>
  );
};

export default MyForm;