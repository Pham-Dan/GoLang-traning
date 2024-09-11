// Filename - components/Home.js

import "bootstrap/dist/css/bootstrap.min.css";
import React, { useContext, useEffect, useState } from "react";
import { Button, Table } from "react-bootstrap";
import { Link, useNavigate } from "react-router-dom";
import { url } from "../Api/url";
import { AuthContext } from "../App";
import { router } from "../route/posts";
import Image from 'react-bootstrap/Image';
import Pagination from 'react-bootstrap/Pagination';
import Dropdown from 'react-bootstrap/Dropdown';

function Home() {

    const [posts, sePPosts] = useState([])
    const [deleteId, setDeleteId] = useState(null)
    const [page, setPage] = useState(1)
    const [total, setTotal] = useState(0)
    const [lastPage, setLastPage] = useState(1)
    const [limit, setLimit] = useState(10)
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
        fetch(url.postGetAll + `?page=${page}&limit=${limit}`, {
            headers: {
                Authorization: token
            },
        })
            .then((res) => {
                if (!res.ok) {
                    throw new Error(`HTTP error: Status ${res.status}`);
                }
                return res.json();
            })
            .then((data) => {
                sePPosts(data.data ?? []);
                setTotal(data.total)
                setLastPage(data.last_page)
            })
            .catch((e) => {
                console.log(e);
            })
    }, [deleteId, page, limit])
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
    const itemsPaginate = []
    for (let number = 1; number <= lastPage; number++) {
        itemsPaginate.push(
            <Pagination.Item key={number} active={number === page} onClick={() => { setPage(number) }}>
                {number}
            </Pagination.Item>
        )

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
            <Dropdown>
                <Dropdown.Toggle variant="success" id="dropdown-basic">
                    Per page {limit}
                </Dropdown.Toggle>

                <Dropdown.Menu>
                    {[5, 10, 15].map((number) => (
                        <Dropdown.Item onClick={() => { setLimit(number) }} href="#">{number}</Dropdown.Item>
                    ))}
                </Dropdown.Menu>
            </Dropdown>
            <Table striped bordered hover responsive className="shadow-sm" style={{ tableLayout: "fixed" }}>
                <caption style={{ captionSide: "top" }}>Total {total} results</caption>
                <thead className="thead-dark">
                    <tr>
                        <th>##</th>
                        <th>Title</th>
                        <th>Content</th>
                        <th>Image</th>
                        <th>User</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {posts.map((item, index) => {
                        return (
                            <tr key={index}>
                                <td>{item.id}</td>
                                <td>{item.title}</td>
                                <td>{item.content}</td>
                                <td>  <Image src={item.image} thumbnail width={100} /></td>
                                <td>{item.User?.name}</td>
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
            <Pagination>
                {itemsPaginate}
            </Pagination>
        </div>
    );
}

export default Home;
