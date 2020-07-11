mvn -f pom.xml dependency:copy-dependencies
curl -O  https://www.antlr.org/download/antlr-4.8-complete.jar
java -jar /Users/didi/PhpstormProjects/c/json-parser/antlr/calc/antlr-4.8-complete.jar



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