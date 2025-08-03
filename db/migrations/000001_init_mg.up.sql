-- Enable foreign keys
PRAGMA foreign_keys = ON;

-- Users table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    display_name TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Chat rooms table
CREATE TABLE rooms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    is_private INTEGER DEFAULT 0, -- 0 = false, 1 = true
    created_by INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by) REFERENCES users (id)
);

-- Room members table
CREATE TABLE room_members (
    room_id INTEGER,
    user_id INTEGER,
    joined_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    role TEXT DEFAULT 'member',
    PRIMARY KEY (room_id, user_id),
    FOREIGN KEY (room_id) REFERENCES chat_rooms (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Messages table
CREATE TABLE messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    room_id INTEGER NOT NULL,
    user_id INTEGER,
    content TEXT NOT NULL,
    sent_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    edited_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (room_id) REFERENCES chat_rooms (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL
);
