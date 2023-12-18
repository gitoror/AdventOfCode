# import os

# # Set the range for i
# start_index = 19
# end_index = 25

# # Read template content from template.go
# template_file_path = "template.go"
# with open(template_file_path, "r") as template_file:
#     template_content = template_file.read()

# # Create folders and files using the template
# for i in range(start_index, end_index + 1):
#     folder_name = str(i)
#     os.makedirs(folder_name, exist_ok=True)

#     # Create empty file and write template content
#     file_name = f"day{i}.go"
#     file_path = os.path.join(folder_name, file_name)
#     with open(file_path, "w") as file:
#         file.write(template_content)

#     print(f"File {file_name} created in folder {folder_name} using the template.")

# print("Folders and files created successfully using the template.")

# for i in range(start_index, end_index + 1):
#     folder_name = "inputs/"+str(i)
#     os.makedirs(folder_name, exist_ok=True)

#     # Create empty file in each folder
#     file_name = f"in.txt"
#     file_path = os.path.join(folder_name, file_name)
#     with open(file_path, "w") as file:
#         pass
#     print(f"Empty file {file_name} created in folder {folder_name}")
#     file_name = f"ex.txt"
#     file_path = os.path.join(folder_name, file_name)
#     with open(file_path, "w") as file:
#         pass
#     print(f"Empty file {file_name} created in folder {folder_name}")

# print("Folders and empty files created successfully.")
