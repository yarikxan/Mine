ALTER TABLE books 
ADD COLUMN isbn VARCHAR(13) UNIQUE,
ADD CONSTRAINT chk_isbn_format CHECK (isbn ~* '^[A-Z0-9]{13}$');