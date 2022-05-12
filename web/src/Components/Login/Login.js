import { useForm } from "react-hook-form"
import { useNavigate } from 'react-router-dom'

export default function Login(){
  const {
    register,
    handleSubmit,
    formState: {errors},
    //watch
  } = useForm()

  const navigate = useNavigate()

  const submit = async data => {
    try {
      const response = await fetch(`${process.env.REACT_APP_API_URL}/admin/login`, {
        method: 'POST',
        body: JSON.stringify(data),
      })
      const res = await response.json()
      localStorage.setItem('token', res.token)

      navigate("/fleek")
    } catch (err) {
      console.log("error on submit", err)
    }
  }
  const style = {
    border: "1px solid grey",
    borderRadius: "5px",
    alignText: "right",
    width: "25%",
    height: "30px"
  }
  return (
    <form action="" onSubmit={handleSubmit(submit)}
    style={{
      display: 'flex',
      flexDirection: "column",
      alignItems: "center",
      height: '100vh',
      justifyContent: "center"
    }}
    >
      <input type="text"  {...register("user", { required: true })} placeholder={"User"} style={
        style
      }/>
      {errors.user?.type === 'required' && (
        <div style={{ color: "red", marginTop: "5px"}}>
          User is a required field
        </div>)}
      <input type="password" {...register("password", { required: true })} placeholder={"Password"} style={
        style
      }/>
      {errors.password?.type === 'required' && (
        <div style={{ color: "red", marginTop: "5px" }}>
          Password is required field
        </div>)}
      <button type="submit" style={{
        border: "2px solid blue",
        background: "white",
        color: "blue",
        width: "50px",
        borderRadius: "5px",
        height: "30px"
      }}>Login</button>
    </form>
  )
}
