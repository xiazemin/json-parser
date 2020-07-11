[INFO] BUILD FAILURE
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 1.383s
[INFO] Finished at: Sat Jul 11 10:26:12 CST 2020
[INFO] Final Memory: 11M/123M
[INFO] ------------------------------------------------------------------------
[ERROR] Failed to execute goal org.apache.maven.plugins:maven-compiler-plugin:2.3.2:compile (default-compile) on project json2xml: Compilation failure: Compilation failure:
[ERROR] /Users/didi/PhpstormProjects/c/json-parser/antlr/json2xml/src/main/java/antlr/json2xml.java:[10,0] 错误: 程序包org.antlr.v4.runtime不存在
[ERROR] /Users/didi/PhpstormProjects/c/json-parser/antlr/json2xml/src/main/java/antlr/json2xml.java:[11,0] 错误: 程序包org.antlr.v4.runtime.tree不存在
[ERROR] /Users/didi/PhpstormProjects/c/json-parser/antlr/json2xml/src/main/java/antlr/JSONBaseListener.java:[12,7] 错误: 类重复: JSONBaseListener
[ERROR] /Users/didi/PhpstormProjects/c/json-parser/antlr/json2xml/src/main/java/antlr/JSONBaseListener.java:[3,27] 错误: 程序包org.antlr.v4.runtime不存在
[ERROR] /Users/didi/PhpstormProjects/c/json-parser/antlr/json2xml/src/main/java/antlr/JSONBaseListener.java:[4,32] 错误: 程序包org.antlr.v4.runtime.tree不存在
[ERROR] /Users/didi/PhpstormProjects/c/json-parser/antlr/json2xml/src/main/java/antlr/JSONBaseListener.java:[5,32] 错误: 程序包org.antlr.v4.runtime.tree不存在
[ERROR] /Users/didi/PhpstormProjects/c/json-parser/antlr/json2xml/src/main/java/antlr/json2xml.java:[42,43] 错误: 无法访问JSONBaseListener
[ERROR] -> [Help 1]
[ERROR]
[ERROR] To see the full stack trace of the errors, re-run Maven with the -e switch.
[ERROR] Re-run Maven using the -X switch to enable full debug logging.
[ERROR]


https://www.cnblogs.com/chywx/p/11563318.html

        <dependency>
            <groupId>org.antlr</groupId>
            <artifactId>antlr4</artifactId>
            <version>4.8.1</version>
            <scope>system</scope>
            <systemPath>${project.basedir}/antlr-4.8-complete.jar</systemPath>
        </dependency>
        <!--dependency>
            <groupId>org.antlr</groupId>
            <artifactId>antlr4</artifactId>
            <version>4.8.1</version>
        </dependency-->

  [ERROR] Failed to execute goal org.apache.maven.plugins:maven-compiler-plugin:2.3.2:compile (default-compile) on project json2xml: Compilation failure: Compilation failure:
  [ERROR] /Users/didi/PhpstormProjects/c/json-parser/antlr/json2xml/src/main/java/antlr/JSONBaseListener.java:[12,7] 错误: 类重复: JSONBaseListener
  [ERROR] /Users/didi/PhpstormProjects/c/json-parser/antlr/json2xml/src/main/java/antlr/json2xml.java:[42,43] 错误: 无法访问JSONBaseListener
  [ERROR] -> [Help 1]

  mvn clean package进行打包
  我们得到了项目的输出，如果有需要的话，就可以复制这个jar文件到其他项目的Classpath中，从而使用HelloWorld类。

  但是，如何让其他的Maven项目直接饮用这个jar呢，还需要一个安装步骤：mvn clean install

  在打包之后，又执行了install。从输出中可以看到该任务将项目输出的jar安装到了Maven本地仓库中，可以打开相应的文件夹看到HelloWorld项目的pom和jar。

  默认打包生成的jar是不能够直接运行的，因为带有main方法的类信息不会添加到manifest中（打开jar文件中的META-INF/MANIFEST.MF文件，将无法看到Main-Class一行）

  为了生成可执行的jar文件，需要借助maven-shade-plugin

  <plugin>
                  <groupId>org.apache.maven.plugins</groupId>
                  <artifactId>maven-shade-plugin</artifactId>
                  <version>1.2.1</version>
                  <executions>
                      <execution>
                          <phase>package</phase>
                          <goals>
                              <goal>shade</goal>
                          </goals>
                          <configuration>
                              <transformers>
                                  <transformer implementation="org.apache.maven.plugins.shade.resource.ManifestResourceTransformer">
                                      <mainClass>com.juvenxu.mvnbook.helloworld.HelloWorld</mainClass>
                                  </transformer>
                              </transformers>
                          </configuration>
                      </execution>
                  </executions>
              </plugin>

  这里配置了mainClass为com.juvenxu.mvnbook.helloworld.HelloWorld，项目打包时会将该信息放到MANIFEST中，再次执行mvn clean install，构建完之后打开target/目录，可以看到hello-world-0.0.1-SNAPSHOT.jar和original-hello-world-0.0.1-SNAPSHOT.jar，前者是带有Main-Class信息的可运行jar，后者是原始的jar
  打开hello-world-0.0.1-SNAPSHOT.jar的META-INF/MANIFEST.MF，可以看到包含这样一样信息：

  Main-Class: com.juvenxu.mvnbook.helloworld.HelloWorld

   java -jar target\hello-world-0.0.1-SNAPSHOT.jar

  控制台输出了Hello Maven

  https://blog.csdn.net/tomato__/article/details/13625497


  <build>
    	<plugins>
    		<plugin>
    			<artifactId>maven-assembly-plugin</artifactId>
    			<configuration>
    				<descriptorRefs><descriptRef>jar-with-dependencies</descriptRef></descriptorRefs>
    				<archive>
    					<manifest>
    						<mainClass>moviespiders.Spider</mainClass>
    					</manifest>
    				</archive>
    			</configuration>
    			<executions>
    				<execution>
    					<id>make-assembly</id>
    					<phase>package</phase>
    					<goals>
    						<goal>assembly</goal>
    					</goals>
    				</execution>
    			</executions>
    		</plugin>
    	</plugins>
    </build>

  mvn clean package
  $java -jar target/json2xml-1.0-SNAPSHOT-jar-with-dependencies.jar

  https://blog.csdn.net/qq_36336003/article/details/80198170

Error: A JNI error has occurred, please check your installation and try again
Exception in thread "main" java.lang.NoClassDefFoundError: org/antlr/v4/runtime/CharStream
	at java.lang.Class.getDeclaredMethods0(Native Method)


  <build>
      <finalName>LeadServer</finalName>   <!-- jar包名前缀，如果没有指定，则会使用{artifactId}作为前缀 -->
      <plugins>
          <plugin>
              <artifactId>maven-assembly-plugin</artifactId>
              <configuration>
                  <archive>
                      <manifest>
                          <mainClass>com.lead.server.Start</mainClass>  <!-- 运行的主函数，根据自己项目的目录修改 -->
                      </manifest>
                  </archive>
                  <descriptorRefs>
                      <descriptorRef>jar-with-dependencies</descriptorRef>  <!-- 表明要添加依赖包 -->
                  </descriptorRefs>
              </configuration>
          </plugin>
      </plugins>
  </build>


  https://www.jianshu.com/p/14bcb17b99e0

  https://www.jianshu.com/p/de69ddebff76