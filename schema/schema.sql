CREATE TABLE example_table (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- -- Add trigger to update modified_at automatically
-- CREATE TRIGGER update_modified_at 
--     AFTER UPDATE ON example_table
--     FOR EACH ROW
--     BEGIN
--         UPDATE example_table SET modified_at = CURRENT_TIMESTAMP
--         WHERE id = OLD.id;
--     END;