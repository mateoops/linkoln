import React, { useState } from 'react';
import { Container, TextField, Button, Typography, Box } from '@mui/material';
import axios from 'axios';


export default function App() {
  const [url, seturl] = useState('');
  const [shortUrl, setShortUrl] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/short/', {
        url,
      });
      setShortUrl(response.data.shortUrl); // Assuming backend returns { shortUrl: "http://short.ly/abc123" }
    } catch (error) {
      console.error('Error shortening URL:', error);
    }
  };

  return (
    <Container maxWidth="sm" style={{ marginTop: '2rem' }}>
      <Typography variant="h4" gutterBottom>
        URL Shortener
      </Typography>
      <Box component="form" onSubmit={handleSubmit} sx={{ mt: 2 }}>
        <TextField
          fullWidth
          label="Enter URL"
          variant="outlined"
          value={url}
          onChange={(e) => seturl(e.target.value)}
          sx={{ mb: 2 }}
        />
        <Button
          type="submit"
          variant="contained"
          color="primary"
          fullWidth
        >
          Shorten
        </Button>
      </Box>
      {shortUrl && (
        <Box sx={{ mt: 3, textAlign: 'center' }}>
          <Typography variant="h6">Shortened URL:</Typography>
          <a href={shortUrl} target="_blank" rel="noopener noreferrer">
            {shortUrl}
          </a>
        </Box>
      )}
    </Container>
  );
};
