import os 
import sys

def generate_report(path):
    file_count = 0
    dir_count = 0
    total_size = 0

    for root, dirs, files in os.walk(path):
        dir_count += len(dirs)
        for file in files:
            file_count += 1
            file_path = os.path.join(root, file)
            if os.path.isfile(file_path):
                total_size += os.path.getsize(file_path)

        return file_count, dir_count, total_size
    
if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python file_report.py <directory_path>")
        sys.exit(1)
    
    target_path = sys.argv[1]

    files, dirs, size = generate_report(target_path)

    print(f"Directory scanned: {target_path}")
    print(f"Total Files: {files}")
    print(f"Total Directories: {dirs}")
    print(f"Total size: {size} bytes")