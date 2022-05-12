import React, { useEffect, useState, Fragment } from "react";
import { set } from "react-hook-form";

export default function MainPage(props){
  const token = localStorage.getItem('token')
  const [apiKeys, setApiKeys] = useState([])
  const fetchData =  async token => {
    try {
      const response = await fetch(`${process.env.REACT_APP_API_URL}/admin/apikeys`, {
        method: 'GET',
        headers: { 'x-admin-token': token},
      })
      const res = await response.json()

      const keys = res.apiKeys.reduce((prev, curr) => {
        return [...prev, curr]
      }, [])

      setApiKeys(keys)
    } catch(error){
      console.log('error on fetch data', error )
    }
  }

  const patchApiKey = async ({data, token}) => {
    try {
       const response = await fetch(`${process.env.REACT_APP_API_URL}/admin/apikeys`, {
        method: 'PATCH',
        body: JSON.stringify(data),
        headers: { 'x-admin-token': token }, 
      })
      const res = await response.json()

    } catch(err) {
      console.log('error on patch api key', err)
    }
  }

  const [shouldFetch, setShouldFetch] = useState(true)
  useEffect(() => {
    if(token && shouldFetch){
      fetchData(token)
      setShouldFetch(!shouldFetch)
    }
  }, [shouldFetch])

  const onClickCreateApiKey = async () => {
    console.log('post on create api key')
    try {
      const response = await fetch(`${process.env.REACT_APP_API_URL}/admin/apikeys`, {
        headers: { 'x-admin-token': token },
        method: 'POST',
      })
      const res = await response.json()
      setApiKeys([...apiKeys, res])
    } catch (err) {
      console.log('error on create api key', err)
    }
  }

  const onClickCheckBox = async apiKey => {
    const token = localStorage.getItem('token')
    await patchApiKey({
      token,
      data: {
        key: apiKey.key,
        enabled: !apiKey.enabled,
      }
    })
    setApiKeys(apiKeys.map(ak => {
      if (ak.key === apiKey.key) {
        return {
          ...ak,
          enabled: !apiKey.enabled,
        }
      }
      return ak
    }))
  }

  return (
    <div>
      <h1>MainPage</h1>
      <></>

      {
        token ? (
          <div>
            <button
              type="button"
              style={buttonStyle}
              onClick={async () => await onClickCreateApiKey()}
            >Create api key</button>
            <></>

            <table style={tableStyle}>
              <></>
              <tbody>
                <tr>
                  <th style={thtrStyle}>Api key</th>
                  <th style={thtrStyle}>Requests made</th>
                  <th style={thtrStyle}>Bytes transfered</th>
                  <th style={thtrStyle}>Enableb</th>
                </tr>

                {apiKeys.map((apiKey, idx) => {
                  return (
                    <tr key={idx}>

                      <th style={thtrStyle}> {apiKey.key}</th>
                      <th style={thtrStyle}> {apiKey.requests}</th>
                      <th style={thtrStyle}> {apiKey.bytesTransfered}</th>
                      <th style={thtrStyle}>
                        <input type="checkbox" checked={apiKey.enabled} onChange={() => onClickCheckBox(apiKey)} />
                      </th>
                    </tr>
                  )
                })}

              </tbody>
            </table>
          </div>
        ) : (<span>please login</span>)
      }
    </div>
  )
}


const tableStyle = {
  margin: "10px",
  borderCollapse: "collapse",
  width: "95%",
}
// styles 
const thtrStyle = {
  border: "1px solid #dddddd",
  textAlign: "left",
  padding: "8px",
}
const buttonStyle = {
  border: "2px solid blue",
  background: "white",
  color: "blue",
  width: "125px",
  borderRadius: "5px",
  height: "30px",
}