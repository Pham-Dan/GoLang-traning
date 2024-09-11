// Filename - components/Create.js

import "bootstrap/dist/css/bootstrap.min.css";
import React, { useState } from "react";
import { Button, Form , Container} from "react-bootstrap";
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
        <Container>
            <h2 className="text-center">Login</h2>
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
                    <Form.Label>Email address</Form.Label>
                    <Form.Control
                        onChange={(e) =>
                            setEmail(e.target.value)
                        }
                        type="text"
                        placeholder="Email"
                        required
                        size="lg"
                    />
                </Form.Group>


                <Form.Group
                    className="mb-3"
                    controlId="formBasicAge"
                >
                    <Form.Label>Password</Form.Label>
                    <Form.Control
                        onChange={(e) =>
                            setPassword(e.target.value)
                        }
                        type="password"
                        placeholder="Password"
                        required
                        size="lg"
                    />
                </Form.Group>


                <Button
                    onClick={(e) => handelSubmit(e)}
                    variant="primary"
                    type="submit"
                    size="lg"
                >
                    Login
                </Button>
                {error && <p className="text-danger text-center">{error}</p>}

            </Form>
        </Container>
    );
}

export default Login;
