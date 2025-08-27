import pandas as pd
import matplotlib.pyplot as plt

# File paths
fast_file = r'resources/Locust_http_requests_FastHttpUser_ecs.csv'
http_file = r'resources/Locust_http_requests_HttpUser_ecs.csv'


def get_weighted_avg_response_times(file_path):
    df = pd.read_csv(file_path)
    # POST
    post_df = df[df['Type'] == 'POST']
    post_total = (post_df['Average Response Time'] * post_df['Request Count']).sum()
    post_count = post_df['Request Count'].sum()
    post_avg = post_total / post_count if post_count > 0 else 0
    # GET
    get_df = df[df['Type'] == 'GET']
    get_total = (get_df['Average Response Time'] * get_df['Request Count']).sum()
    get_count = get_df['Request Count'].sum()
    get_avg = get_total / get_count if get_count > 0 else 0
    return post_avg, get_avg


fast_post, fast_get = get_weighted_avg_response_times(fast_file)
http_post, http_get = get_weighted_avg_response_times(http_file)

# Prepare data for plotting
labels = ['FastHttpUser', 'HttpUser']
post_times = [fast_post, http_post]
get_times = [fast_get, http_get]

x = range(len(labels))
width = 0.35

fig, ax = plt.subplots()
ax.bar(x, post_times, width, label='POST', color='skyblue')
ax.bar([i + width for i in x], get_times, width, label='GET', color='coral')

ax.set_ylabel('Weighted Avg Response Time (ms)')
ax.set_title('Weighted Avg Response Time by Request Type')
ax.set_xticks([i + width/2 for i in x])
ax.set_xticklabels(labels)
ax.legend(loc='upper right')

plt.tight_layout()
plt.savefig('resources/response_time_comparison_assignment_4b.png')
plt.close()
print("Graph saved to resources/response_time_comparison_assignment_4b.png")