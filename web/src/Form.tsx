import { useEffect, useState } from 'react';

interface UserSuggestion {
    id: number;
    first_name: string;
  }

function Form() {

    const [user, setUser] = useState('');
    const [message, setMessage] = useState('');
    const [suggestions, setSuggestions] =useState<UserSuggestion[]>([])

    useEffect(() => {
        const fetchSuggestions = async () => {
            if (user.trim() === '') {
                setSuggestions([]);
                return;
            }

            try {
                const response = await fetch(`http://localhost:8000/v1/users?search=${user}`);
                if (response.ok) {
                    const data = await response.json();
                    setSuggestions(data);
                    console.log(data)
                  } else {
                    console.error('Error fetching user suggestions');
                  }
            } catch (error) {
                console.error('Error:', error);
            }
        }
        fetchSuggestions();
    }, [user]);

  return (
    <div className='bg-slate-500'>        
        <input
            type="text"
            className='input input-bordered w-full max-w-xs mb-4'
            id="user"
            value={user}
            onChange={(event) => setUser(event.target.value)}
        />
        {suggestions.length > 0 && (
          <ul>
            {suggestions.map((suggestion) => (
              <li key={suggestion.id} onClick={() => setUser(suggestion.first_name)}>
                {suggestion.first_name}
              </li>
            ))}
          </ul>
        )}
    </div>
  )
}

export default Form