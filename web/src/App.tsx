import { useRef } from "react";

function App() {

  const user = useRef(null);
  const message = useRef(null);
  const currentUser = "";

  const addMessage = () => {
    fetch('http://localhost:8080/v1/posts',{
      method: 'POST',
      body: JSON.stringify({
        author_id: currentUser,
        user_id: user,
        message: message
      }),
      headers: {
        'Content-type': 'application/json'
      }
    })

      .then((response) => response.json())
      .then((data) => {
        console.log(data)
      })
      .catch((err) => {
        console.log(err.message)
      })
  }

  return (
    <>
    <div className="flex justify-center pt-5">
    <div className="">
      <div className="pb-3">
      <input type="text" placeholder="Find user" className="input input-bordered w-full max-w-xs mb-4" ref={user} />
      <textarea className="textarea textarea-bordered textarea-md w-full max-w-xs" placeholder="Message" ref={message}></textarea>
      </div>
      <button className="btn btn-primary" onClick={addMessage}>Add Message</button>
      </div>
      </div>
    </>
  )
}

export default App
