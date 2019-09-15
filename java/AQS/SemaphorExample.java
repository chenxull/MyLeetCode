import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class SemaphorExample{

    private static final int threadCount = 500;
    public static void main(String[] args)throws InterruptedException{
        ExecutorService threadPool = Executors.newFixedThreadPool(300);

        final Semaphore semaphore = new Semaphore(20);

        for (int i=0;i < threadCount ;i++){
            final int threadnum = i;
            threadPool.execute(() ->{
                try{
                    semaphore.acquire();
                    test(threadnum);
                    semaphore.release();
                }catch(InterruptedException e){
                    e.printStackTrace();
                }
            });
        }
        threadPool.shutdown();
        System.out.println("finish");
    }

    public static void test(int threadnum)throws InterruptedException{
        Thread.sleep(1000);// 模拟请求的耗时操作
        System.out.println("threadnum:" + threadnum);
        Thread.sleep(1000);// 模拟请求的耗时操作
 
    }
}