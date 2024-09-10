// Filename - Edit.js
import React, { useContext, useEffect, useState } from "react";
import { Button, Form } from "react-bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "../App";
import { url } from "../Api/url";

function Edit() {

    const [post, setPost] = useState();
    const [id, setid] = useState(localStorage.getItem("id"));

    // Used for navigation with logic in javascript
    let history = useNavigate();
    const token = useContext(AuthContext);
    useEffect(() => {
        if (!token) {
            history('/login')
        }
    }, [token])

    const handelSubmit = (e) => {
        // Preventing from reload
        e.preventDefault();
        if (post.title == "" || post.content == "") {
            alert("invalid input");
            return;
        }
        fetch(url.postUpdate(id), {
            method: "put",
            headers: {
                "Content-Type": "application/json",
                Authorization: token
            },
            body: JSON.stringify(post)

        })
            .then((res) => {
                if (res.ok) {
                    history('/')
                }
            })

            .catch((e) => {
                console.log(e);
            })

    };
    useEffect(() => {
        fetch(url.postGet(id), {
            headers: {
                Authorization: token
            }
        })
            .then((res) => {
                if (!res.ok) {
                    throw new Error(`HTTP error: Status ${res.status}`);
                }
                return res.json();
            })
            .then((data) => {
                console.log(post);
                setPost(data.data);
            })
            .catch((e) => {
                console.log(e);
            })
    }, [])

    return (
        <div>
            <Form
                className="d-grid gap-2"
                style={{ margin: "5rem" }}
            >
                {/* setting a name from the 
                    input textfiled */}
                <Form.Group
                    className="mb-3"
                    controlId="formBasicEmail"
                >
                    <Form.Control
                        value={post?.title}
                        onChange={(e) =>
                            setPost({ ...post, title: e.target.value })
                        }
                        type="text"
                        placeholder="Title"
                    />
                </Form.Group>

                {/* setting a age from the input textfiled */}
                <Form.Group
                    className="mb-3"
                    controlId="formBasicPassword"
                >
                    <Form.Control
                        value={post?.content}
                        onChange={(e) =>
                            setPost({ ...post, content: e.target.value })

                        }
                        type="text"
                        placeholder="Content"
                    />
                </Form.Group>

                {/* Hadinling an onclick event 
                    running an edit logic */}
                <Button
                    onClick={(e) => handelSubmit(e)}
                    variant="primary"
                    type="submit"
                    size="lg"
                >
                    Update
                </Button>

                {/* Redirecting to main page after editing */}
                <Link className="d-grid gap-2" to="/">
                    <Button variant="warning" size="lg">
                        Home
                    </Button>
                </Link>
            </Form>
        </div>
    );
}

export default Edit;
