// app/posts/page.tsx - Server Component 示例
interface Post {
    id: number;
    title: string;
    body: string;
  }
  
  export default async function PostsPage() {
    // 服务器组件中可以直接访问所有环境变量
    const apiUrl = process.env.API_BASE_URL || 'https://jsonplaceholder.typicode.com';
    const apiKey = process.env.API_SECRET_KEY; // 这是私密信息，不会暴露给客户端
    
    console.log('🔐 使用 API Key:', apiKey ? '已配置' : '未配置');
    console.log('🌐 API 地址:', apiUrl);
    
    try {
      // 在服务器端获取数据
      const response = await fetch(`${apiUrl}/posts?_limit=5`, {
        headers: {
          'Content-Type': 'application/json',
          // 如果有 API Key，添加到请求头（不会暴露给客户端）
          ...(apiKey && { 'Authorization': `Bearer ${apiKey}` }),
        },
        // 添加缓存策略（Next.js 特有）
        next: { revalidate: 60 } // 每60秒重新验证
      });
      
      if (!response.ok) {
        throw new Error(`API 请求失败: ${response.status}`);
      }
      
      const posts: Post[] = await response.json();
      
      return (
        <div className="container mx-auto p-4">
          <h1 className="text-3xl font-bold mb-6">文章列表</h1>
          <p className="text-sm text-gray-500 mb-4">
            从 {apiUrl} 获取（服务器端渲染）
          </p>
          
          <div className="space-y-4">
            {posts.map((post) => (
              <article key={post.id} className="border rounded-lg p-4 shadow-sm">
                <h2 className="text-xl font-semibold mb-2">{post.title}</h2>
                <p className="text-gray-600">{post.body}</p>
              </article>
            ))}
          </div>
          
          <div className="mt-6 p-4 bg-blue-50 rounded-lg">
            <p className="text-sm">
              💡 <strong>提示：</strong>这个页面在服务器端渲染，
              环境变量 <code className="bg-gray-200 px-1 rounded">API_SECRET_KEY</code> 永远不会暴露给浏览器
            </p>
          </div>
        </div>
      );
    } catch (error) {
      console.error('❌ 获取文章失败:', error);
      return (
        <div className="container mx-auto p-4">
          <h1 className="text-3xl font-bold mb-4 text-red-600">加载失败</h1>
          <p className="text-gray-600">
            无法获取文章列表，请检查服务器日志
          </p>
        </div>
      );
    }
  }