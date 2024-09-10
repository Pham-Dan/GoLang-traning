// Filename - components/Home.js

import "bootstrap/dist/css/bootstrap.min.css";
import React, { useContext, useEffect, useState } from "react";
import { Button, Table } from "react-bootstrap";
import { Link, useNavigate } from "react-router-dom";
import { url } from "../Api/url";
import { AuthContext } from "../App";
import { router } from "../route/posts";
import Image from 'react-bootstrap/Image';

function Home() {

    const [posts, sePPosts] = useState([])
    const [deleteId, setDeleteId] = useState(null)
    let history = useNavigate();
    const token = useContext(AuthContext);


    // Function to set the ID, Name, and Age in local storage
    function setID(id) {
        localStorage.setItem("id", id);
    }

    // Function to delete an entry
    function deleted(id) {
        if (window.confirm("Delete post?")) {
            fetch(url.postDelete(id), {
                method: "DELETE",
                headers: {
                    Authorization: token
                }
            })
                .then((res) => {
                    if (res.ok) {
                        console.log('ok');
                        setDeleteId(id)
                        history('/')
                    }
                })
                .catch((e) => {
                    console.log(e);
                })

        }

    }
    useEffect(() => {
        if (!token) {
            history('/login')
        }
    }, [token])
    useEffect(() => {
        fetch(url.postGetAll, {
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
                console.log(data);
                sePPosts(data.data ?? []);
            })
            .catch((e) => {
                console.log(e);
            })
    }, [deleteId])
    const exportCsv = () => {
        fetch(url.postExportCsv, {
            headers: {
                Authorization: token,
                'Content-Type': 'text/csv',
            }
        })
            .then((res) => {
                if (res.ok) {
                    return res.blob()
                }
            })
            .then((blob) => {
                // Create a URL for the CSV file
                const url = window.URL.createObjectURL(new Blob([blob], { type: 'text/csv' }));
                const link = document.createElement('a');
                link.href = url;
                link.setAttribute('download', `posts-${Date.now()}.csv`); // Specify the file name
                document.body.appendChild(link);
                link.click();
                link.remove();
            })
            .catch((e) => {
                console.log(e);
            })
    }
    return (
        <div style={{ margin: "2rem" }}>
            <h2 className="text-center mb-4">Posts Management</h2>
            <div className="mb-4 d-flex justify-content-end gap-5">
                <Button onClick={exportCsv} >ExportCsv</Button>
                <Link to={router.postCreate}>
                    <Button variant="success" size="lg">
                        Create Post
                    </Button>
                </Link>
            </div>
            <Table striped bordered hover responsive className="shadow-sm">
                <thead className="thead-dark">
                    <tr>
                        <th>##</th>
                        <th>Title</th>
                        <th>Content</th>
                        <th>Image</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {posts.map((item, index) => {
                        return (
                            <tr key={index}>
                                <td>{++index}</td>
                                <td>{item.title}</td>
                                <td>{item.content}</td>
                                <td>  <Image src={item.image} thumbnail width={100} /></td>
                                <td>
                                    <Link to={router.postEdit}>
                                        <Button
                                            onClick={() => setID(item.id)}
                                            variant="info"
                                            className="me-2"
                                        >
                                            Update
                                        </Button>
                                    </Link>
                                    <Button
                                        onClick={() => deleted(item.id)}
                                        variant="danger"
                                    >
                                        Delete
                                    </Button>
                                </td>
                            </tr>
                        );
                    })}
                </tbody>
            </Table>
        </div>
    );
}

export default Home;
