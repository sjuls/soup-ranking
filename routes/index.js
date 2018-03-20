const express = require('express');
const router = express.Router();

const rankings = require('./ranking').currentRankings

router.get('/_status', (req, res, next) => res.json({ title: 'I am ALIVE!' }));

router.post('/score', (req, res, next) => {
  rankings.push(req.body);
  res.status(200).send();
});

router.get('/score', (req, res, next) => res.json(rankings));

module.exports = router;
