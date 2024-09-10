// Filename - components/Create.js

import React, { useContext, useEffect, useState } from "react";
import { Button, Form } from "react-bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import array from "./array";
import { v4 as uuid } from "uuid";
import { Link, useNavigate } from "react-router-dom";
import { url } from "../Api/url";
import { AuthContext } from "../App";

function Create() {

    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");
    const [image, setImage] = useState(null);
    const [imagePre, setImagePre] = useState("");
    const [error, setError] = useState("");

    // Using useNavigation for redirecting to pages
    let history = useNavigate();
    const token = useContext(AuthContext);
    
    // Function for creating a post/entry
    const handelSubmit = (e) => {
        e.preventDefault(); // Prevent reload

        const formData = new FormData();
        formData.append("title", title);
        formData.append("content", content);
        formData.append("image", image);
        fetch(url.postCreate, {
            method: "post",
            headers: {
                // "Content-Type": "multipart/form-data",
                Authorization: token
            },
            body: formData

        })
            .then((res) => {
                if (res.ok) {
                    history('/')
                }
                return res.json()
            })
            .then((data) => {
                if (data.error) {
                    setError(data.error)
                }
            })

            .catch((e) => {
                alert(e)
            })

        // Redirecting to home page after creation done
    };

    useEffect(() => {
       if (image instanceof Blob) {
        let url = URL.createObjectURL(image)
       setImagePre(url)
       }
    },[image])

    useEffect(() => {
        if (!token) {
            history('/login')
        }
    },[token])

    return (
        <div>
            <Form
                className="d-grid gap-2"
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
                            setTitle(e.target.value)
                        }
                        type="text"
                        placeholder="Title"
                        required
                    />
                    {error.Title && <p className="text-danger" style={{ "textAlign": "left" }}>{error.Title}</p>}
                </Form.Group>

                {/* Fetching a value from input textfirld in
                    a setContent using usestate*/}
                <Form.Group
                    className="mb-3"
                    controlId="formBasicAge"
                >
                    <Form.Control
                        onChange={(e) =>
                            setContent(e.target.value)
                        }
                        type="text"
                        placeholder="Content"
                        required
                    />
                    {error.Content && <p className="text-danger" style={{ "textAlign": "left" }}>{error.Content}</p>}
                </Form.Group>
                <Form.Group
                    className="mb-3"
                    controlId="formBasicAge"
                >
                    <Form.Control
                        onChange={(e) =>
                            setImage(e.target.files[0])
                        }
                        type="file"
                        placeholder="Image"
                        required
                    />
                </Form.Group>
                {imagePre && <img src={imagePre} height={200}/>}

                {/* handing a onclick event in button for
                    firing a function */}
                <Button
                    onClick={(e) => handelSubmit(e)}
                    variant="primary"
                    type="submit"
                >
                    Submit
                </Button>

                {/* Redirecting back to home page */}
                <Link className="d-grid gap-2" to="/">
                    <Button variant="info" size="lg">
                        Home
                    </Button>
                </Link>
            </Form>
        </div>
    );
}

export default Create;
