# Synth player
This very simple WIP-script plays notes using the [github.com/200sc/klangsynthese](https://github.com/200sc/klangsynthese) synthesizer.
The input is a text file with each line in the following format:

```
[Note start time] [Note duration] [Note A-G with Octave 0-8 and optional Sharp s]
```

For an example see the `notes.txt` file.

### Usage example
- Teach ChatGPT the format
- Ask for a sequence, for example in a certain key
- Add the sequence to the `notes.txt` file
- Play!