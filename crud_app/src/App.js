// Filename - App.js

import React, { createContext, useEffect, useState } from "react";
import { Button, Container } from "react-bootstrap";
import {
    Route,
    BrowserRouter as Router,
    Routes,
} from "react-router-dom";
import "./App.css";
import Create from "./components/Create";
import Edit from "./components/Edit";
import Home from "./components/Home";
import Login from "./components/Login";
import { router } from "./route/posts";
import { url } from "./Api/url";
import Navbar from 'react-bootstrap/Navbar';

export const AuthContext = createContext()
function App() {


    const [token, setToken] = useState(localStorage.getItem("token"));
    const [user, setUser] = useState(localStorage.getItem("user"));

    useEffect(() => {
        if (token) {
            localStorage.setItem("token", token)
            fetch(url.profile, {
                headers: {
                    Authorization: token
                }
            })
                .then((res) => {
                    if (!res.ok) {
                        if (res.status == 401) {
                            localStorage.removeItem("token");
                            setToken(null)
                        }
                        throw new Error(`HTTP error: Status ${res.status}`);
                    }
                    return res.json();
                })
                .then((data) => {
                    setUser(data);
                })
                .catch((e) => {
                    console.log(e);
                })
        }

    }, [token])

    const logout = () => {

        if (window.confirm("logout?")) {
            localStorage.removeItem("token");
            setToken(null)
        }
    }
    return (
        <div className="App">
            <Navbar bg="dark" data-bs-theme="dark">
                <Container>
                    <Navbar.Brand href="#home" className="text-white">App Posts</Navbar.Brand>
                </Container>
            </Navbar>
            <Container>
                {token && <div className="d-flex justify-content-end gap-4 align-items-center mt-5">
                    <h3>{user?.name} - {user?.email}</h3>
                    <Button variant="warning" size="lg" onClick={logout}>
                        Log out
                    </Button>
                </div>}
                <Router>
                    <AuthContext.Provider value={token}>
                        <Routes>
                            <Route path="/" element={<Home />} />
                            <Route path="/login" element={<Login setToken={setToken} />} />
                            <Route
                                path={router.postCreate}
                                element={<Create />}
                            />
                            <Route
                                path={router.postEdit}
                                element={<Edit />}
                            />
                        </Routes>
                    </AuthContext.Provider>
                </Router>
            </Container>
        </div>
    );
}

export default App;

