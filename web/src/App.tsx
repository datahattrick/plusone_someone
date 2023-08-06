import { useRef } from "react";
import Form from "./Form";

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
      <Form />
      </div>
      </div>
    </>
  )
}

export default App
