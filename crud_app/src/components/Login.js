// Filename - components/Create.js

import "bootstrap/dist/css/bootstrap.min.css";
import React, { useState } from "react";
import { Button, Form } from "react-bootstrap";
import { useNavigate } from "react-router-dom";
import { url } from "../Api/url";

function Login({setToken}) {
    let history = useNavigate();

    const [email, setEmail] = useState();
    const [password, setPassword] = useState();
    const [error, setError] = useState();

    // Function for creating a post/entry
    const handelSubmit = (e) => {
        e.preventDefault(); // Prevent reload

        fetch(url.login, {
            method: "post",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                email, password
            })

        })
            .then((res) => {
                if (!res.ok) {
                    throw new Error("These credentials do not match our records!");
                }
                return res.json();
            }).then((data) => {
                setToken(data.accessToken)
                history("/")
            })

            .catch((e) => {
                setError(e.message)
            })

        // Redirecting to home page after creation done
    };

    return (
        <div>
            <h2>Login</h2>
            <Form
                className="d-grid  gap-2"
                style={{ margin: "5rem" }}
            >
                {/* Fetching a value from input textfirld 
                    in a setTitle using usestate*/}
                <Form.Group
                    className="mb-3"
                    controlId="formBasicName"
                >
                    <Form.Control
                        onChange={(e) =>
                            setEmail(e.target.value)
                        }
                        type="text"
                        placeholder="Email"
                        required
                    />
                </Form.Group>


                <Form.Group
                    className="mb-3"
                    controlId="formBasicAge"
                >
                    <Form.Control
                        onChange={(e) =>
                            setPassword(e.target.value)
                        }
                        type="password"
                        placeholder="Password"
                        required
                    />
                </Form.Group>


                <Button
                    onClick={(e) => handelSubmit(e)}
                    variant="primary"
                    type="submit"
                >
                    Login
                </Button>
                {error && <p className="text-danger">{error}</p>}

            </Form>
        </div>
    );
}

export default Login;
