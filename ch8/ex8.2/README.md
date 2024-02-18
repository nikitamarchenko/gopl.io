# FTP

To much for 1 day exercise. Need huge optimize and refactor.

## Basic spec

5.1.  MINIMUM IMPLEMENTATION

In order to make FTP workable without needless error messages, the
following minimum implementation is required for all servers:

    TYPE - ASCII Non-print
    MODE - Stream
    STRUCTURE - File, Record
    COMMANDS - USER, QUIT, PORT,
            TYPE, MODE, STRU,
                for the default values
            RETR, STOR,
            NOOP.

The default values for transfer parameters are:

    TYPE - ASCII Non-print
    MODE - Stream
    STRU - File

All hosts must accept the above as the standard defaults.

5.3.1.  FTP COMMANDS

The following are the FTP commands:

USER <SP> <username> <CRLF>
CWD  <SP> <pathname> <CRLF>
QUIT <CRLF>
PORT <SP> <host-port> <CRLF>
TYPE <SP> <type-code> <CRLF>
STRU <SP> <structure-code> <CRLF>
MODE <SP> <mode-code> <CRLF>
RETR <SP> <pathname> <CRLF>
STOR <SP> <pathname> <CRLF>
PWD  <CRLF>
LIST [<SP> <pathname>] <CRLF>
NOOP <CRLF>


## Basic Hello

```
220 Hello
USER username
331 User ok, Waiting for the password.
PASS password
230 Login Ok.
SYST
215 UNIX Type: L8
FEAT
211-Extensions supported:
PASV
UTF8
211 End.
TYPE I
200 TYPE is now 8-bit binary
QUIT
221 Logout.
```

## List

```
PORT 127,0,0,1,176,185
200 connection accepted
LIST
150 Accepted data connection
226 84 matches total
```

LIST [<SP> <pathname>]
    125, 150
        226, 250
        425, 426, 451
    450
    500, 501, 502, 421, 530

125 Data connection already open; transfer starting.
150 Accepted data connection
    226 84 matches total
    250 Requested file action okay, completed.

    425 Can't open data connection.
    426 Connection closed; transfer aborted.
    451 Requested action aborted: local error in processing.
450 Requested file action not taken. //File unavailable (e.g., file busy).


## Get

RETR
    125, 150
        (110)
        226, 250
        425, 426, 451
    450, 550
    500, 501, 421, 530
