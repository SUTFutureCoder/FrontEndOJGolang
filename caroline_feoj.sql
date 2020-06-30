/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : localhost
 Source Database       : caroline_feoj

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : utf-8

 Date: 06/30/2020 17:44:39 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `lab`
-- ----------------------------
DROP TABLE IF EXISTS `lab`;
CREATE TABLE `lab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `lab_name` varchar(128) NOT NULL DEFAULT '' COMMENT '实验室名称',
  `lab_desc` text NOT NULL COMMENT '实验室描述',
  `lab_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '实验室类型',
  `lab_sample` text NOT NULL COMMENT '实验室样例或地址',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '实验室状态',
  `creator` char(32) NOT NULL DEFAULT '' COMMENT '创建人',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `lab_type` (`lab_type`),
  KEY `status` (`status`),
  KEY `creator` (`creator`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='实验室表';

-- ----------------------------
--  Records of `lab`
-- ----------------------------
BEGIN;
INSERT INTO `lab` VALUES ('18', 'FirstLab', 'FirstLab', '0', '<html>\n<body>\n\n<h1>我的第一个标题</h1>\n\n<p>我的第一个段落。</p>\n\n</body>\n</html>', '1', 'CaveJohson', '487', '0'), ('19', 'FirstLab', '', '0', '', '1', 'CaveJohson', '911', '0'), ('20', 'FirstLab', 'FirstLab', '0', '<html>\n<body>\n\n<h1>我的第一个标题</h1>\n\n<p>我的第一个段落。</p>\n\n</body>\n</html>', '1', 'CaveJohson', '1593262590826', '0'), ('21', '4D时空序列复杂实验室', 'CASE1 在最初的默认值，需要指定#input的textarea输入为1\\n2\\n3\\n4\\n5\\n，对应的#output的textarea输出值为输入的2倍，即2\\n4\\n6\\n8\\n10\nCASE2 要求#input和#output在同一行\nCASE3 接下来在点击#submit按钮之后，每隔300秒对动态变化的#input进行响应，计算方法为#input每行加1，并输出至#output的textarea中。例如测试2\\n4\\n5\\n6\\n8,输出值为3\\n5\\n6\\n7\\n9\n\nINPUT\n\nOUTPUT\n2\n3\n4\n8\n10\n12\n\nINPUT\ndocument.getElementById(\"input\").getElementsByTagName(\"textarea\")[0].innerHTML = \"2\\n4\\n5\\n6\\n8\"\nOUTPUT\n\"3\n5\n6\n7\n9\"\n', '4', '<!DOCTYPE html>\n\n<!--\nINPUT\n\nOUTPUT\n2\n3\n4\n8\n10\n12\n\nINPUT\ndocument.getElementById(\"input\").getElementsByTagName(\"textarea\")[0].innerHTML = \"2\\n4\\n5\\n6\\n8\"\nOUTPUT\n\"3\n5\n6\n7\n9\"\n\n-->\n\n\n\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Title</title>\n</head>\n<style>\n    .inline {\n        display: inline-block;\n    }\n    #submit_div {\n        width: 100%;\n    }\n</style>\n<script>\n    function doSubmit() {\n        setInterval(changeOutput, 300)\n    }\n\n    function changeOutput() {\n        let input = document.getElementById(\"input\").getElementsByTagName(\"textarea\")[0].innerHTML\n        let inputs = input.split(\"\\n\")\n        let outputs = []\n        for (let i = 0; i < inputs.length; i++) {\n            if (inputs[i] == NaN) {\n                continue\n            }\n            outputs[i] = parseInt(inputs[i]) + 1\n        }\n        let output = outputs.join(\"\\n\")\n        document.getElementById(\"output\").getElementsByTagName(\"textarea\")[0].innerHTML = output\n    }\n</script>\n<body>\n        <div id=\"input\" class=\"inline\">\n<textarea cols=\"5\" rows=\"8\">\n1\n2\n3\n4\n5\n6</textarea>\n        </div>\n\n        <div id=\"output\" class=\"inline\">\n<textarea cols=\"5\" rows=\"8\" disabled=\"disabled\">\n2\n3\n4\n8\n10\n12</textarea>\n        </div>\n\n    <div id=\"submit_div\">\n        <button id=\"submit\" onclick=\"doSubmit()\">运行</button>\n    </div>\n\n</body>\n</html>', '1', 'CaveJohson', '1593505928813', '0');
COMMIT;

-- ----------------------------
--  Table structure for `lab_submit`
-- ----------------------------
DROP TABLE IF EXISTS `lab_submit`;
CREATE TABLE `lab_submit` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `lab_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '实验室id',
  `submit_data` text NOT NULL COMMENT '提交内容',
  `submit_result` text NOT NULL COMMENT '提交结果',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '实验室状态',
  `creator` char(32) NOT NULL DEFAULT '' COMMENT '创建人',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `lab_id` (`lab_id`),
  KEY `status` (`status`),
  KEY `creator` (`creator`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COMMENT='提交表';

-- ----------------------------
--  Records of `lab_submit`
-- ----------------------------
BEGIN;
INSERT INTO `lab_submit` VALUES ('18', '18', '<html>\n<body>\n<h1>我的第一个标题</h1>\n<p>我的第一个段落。</p>\n</body>\n</html>', '', '1', 'Chell01', '1593168274267', '0'), ('19', '21', '<!DOCTYPE html>\n\n<!--\nCASE1 在最初的默认值，需要指定#input的textarea输入为1\\n2\\n3\\n4\\n5\\n，对应的#output的textarea输出值为输入的2倍，即2\\n4\\n6\\n8\\n10\nCASE2 要求#input和#output在同一行\nCASE3 接下来在点击#submit按钮之后，每隔300秒对动态变化的#input进行响应，计算方法为#input每行加1，并输出至#output的textarea中。例如测试2\\n4\\n5\\n6\\n8,输出值为3\\n5\\n6\\n7\\n9\n\nINPUT\n\nOUTPUT\n2\n3\n4\n8\n10\n12\n\nINPUT\ndocument.getElementById(\"input\").getElementsByTagName(\"textarea\")[0].innerHTML = \"2\\n4\\n5\\n6\\n8\"\nOUTPUT\n\"3\n5\n6\n7\n9\"\n\n-->\n\n\n\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Title</title>\n</head>\n<style>\n    .inline {\n        display: inline-block;\n    }\n    #submit_div {\n        width: 100%;\n    }\n</style>\n<script>\n    function doSubmit() {\n        setInterval(changeOutput, 300)\n    }\n\n    function changeOutput() {\n        let input = document.getElementById(\"input\").getElementsByTagName(\"textarea\")[0].innerHTML\n        let inputs = input.split(\"\\n\")\n        let outputs = []\n        for (let i = 0; i < inputs.length; i++) {\n            if (inputs[i] == NaN) {\n                continue\n            }\n            outputs[i] = parseInt(inputs[i]) + 1\n        }\n        let output = outputs.join(\"\\n\")\n        document.getElementById(\"output\").getElementsByTagName(\"textarea\")[0].innerHTML = output\n    }\n</script>\n<body>\n        <div id=\"input\" class=\"inline\">\n<textarea cols=\"5\" rows=\"8\">\n1\n2\n3\n4\n5\n6</textarea>\n        </div>\n\n        <div id=\"output\" class=\"inline\">\n<textarea cols=\"5\" rows=\"8\" disabled=\"disabled\">\n2\n3\n4\n8\n10\n12</textarea>\n        </div>\n\n    <div id=\"submit_div\">\n        <button id=\"submit\" onclick=\"doSubmit()\">运行</button>\n    </div>\n\n</body>\n</html>', '', '1', 'Chell01', '1593508029634', '0');
COMMIT;

-- ----------------------------
--  Table structure for `lab_testcase`
-- ----------------------------
DROP TABLE IF EXISTS `lab_testcase`;
CREATE TABLE `lab_testcase` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `testcase_desc` text NOT NULL COMMENT '测试用例描述',
  `testcase_code` text NOT NULL COMMENT '测试用例代码',
  `input` text NOT NULL COMMENT '测试用例输入',
  `output` text NOT NULL COMMENT '测试用例输出',
  `time_limit` int(10) NOT NULL COMMENT '测试用例时间限制',
  `mem_limit` int(10) NOT NULL COMMENT '测试用例内存限制',
  `wait_before` int(10) NOT NULL DEFAULT '0' COMMENT '用例执行前等待',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '实验室状态',
  `creator` char(32) NOT NULL DEFAULT '' COMMENT '创建人',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `status` (`status`),
  KEY `creator` (`creator`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COMMENT='测试用例表';

-- ----------------------------
--  Records of `lab_testcase`
-- ----------------------------
BEGIN;
INSERT INTO `lab_testcase` VALUES ('18', '用于测试的第一个测试用例', 'document.getElementsByTagName(\"H1\")[0].innerText', '', '我的第一个标题', '0', '0', '0', '1', 'CaveJohson', '579', '0'), ('19', '用于测试的第一个测试用例', 'document.getElementsByTagName(\"p\")[0].innerText', '', '我的第一个段落。', '100', '100', '0', '1', 'CaveJohson', '1593167166947', '0'), ('20', '第一个实验室的测试用例', 'if (document.getElementsByTagName(\"p\").length == 0){ \"缺少段落字段\"}', '', '', '100', '100', '0', '1', 'CaveJohson', '1593167305803', '0'), ('21', '第一个实验室的测试用例', 'if (document.getElementsByTagName(\"p\").length == 0){ \"缺少段落字段\"}', '', '', '100', '100', '0', '0', 'CaveJohson', '1593167306638', '0'), ('22', 'inline测试用例', 'document.defaultView.getComputedStyle(document.getElementsByClassName(\"inline\")[0]).display;', '', 'inline-block', '0', '0', '0', '1', 'CaveJohson', '1593506839989', '0'), ('23', '在最初的默认值，需要指定#input的textarea输入为1\\n2\\n3\\n4\\n5\\n，对应的#output的textarea输出值为输入的2倍，即2\\n4\\n6\\n8\\n10', 'document.getElementById(\"input\").getElementsByTagName(\"textarea\")[0].innerHTML', '', '1\n2\n3\n4\n5\n6', '0', '0', '0', '1', 'CaveJohson', '1593507479378', '0'), ('24', '在最初的默认值，需要指定#input的textarea输入为1\\n2\\n3\\n4\\n5\\n，对应的#output的textarea输出值为输入的2倍，即2\\n4\\n6\\n8\\n10', 'document.getElementById(\"output\").getElementsByTagName(\"textarea\")[0].innerHTML', '', '2\n3\n4\n8\n10\n12', '0', '0', '0', '1', 'CaveJohson', '1593507499325', '0'), ('25', '在200毫秒时，应该是原始值', 'if (typeof clicked==\"undefined\") {document.getElementById(\"submit\").click(); clicked = true;} document.getElementById(\"output\").getElementsByTagName(\"textarea\")[0].innerHTML', '', '2\n3\n4\n8\n10\n12', '0', '0', '200', '1', 'CaveJohson', '1593507764518', '0'), ('26', '在400毫秒时，应该是输入值加1', 'if (typeof clicked==\"undefined\") {document.getElementById(\"submit\").click(); clicked = true;} document.getElementById(\"output\").getElementsByTagName(\"textarea\")[0].innerHTML', '', '2\n3\n4\n5\n6\n7', '0', '0', '400', '1', 'CaveJohson', '1593507791163', '0'), ('27', '在400毫秒时，应该是能够响应输入值变化后加1', 'if (typeof clicked==\"undefined\") {document.getElementById(\"submit\").click(); clicked = true;} document.getElementById(\"input\").getElementsByTagName(\"textarea\")[0].innerHTML=\"30\\n40\\n50\";document.getElementById(\"output\").getElementsByTagName(\"textarea\")[0].innerHTML;', '', '31\n41\n51', '0', '0', '400', '1', 'CaveJohson', '1593507993386', '0');
COMMIT;

-- ----------------------------
--  Table structure for `lab_testcase_map`
-- ----------------------------
DROP TABLE IF EXISTS `lab_testcase_map`;
CREATE TABLE `lab_testcase_map` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `lab_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '实验室id',
  `testcase_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '测试用例id',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '实验室状态',
  `creator` char(32) NOT NULL DEFAULT '' COMMENT '创建人',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `lab_id` (`lab_id`),
  KEY `testcase_id` (`testcase_id`),
  KEY `status` (`status`),
  KEY `creator` (`creator`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COMMENT='实验室-测试用例关联表';

-- ----------------------------
--  Records of `lab_testcase_map`
-- ----------------------------
BEGIN;
INSERT INTO `lab_testcase_map` VALUES ('18', '18', '18', '1', 'CaveJohson', '579', '0'), ('19', '18', '19', '1', 'CaveJohson', '1593167166947', '0'), ('20', '18', '20', '1', 'CaveJohson', '1593167305803', '0'), ('21', '18', '21', '0', 'CaveJohson', '1593167306638', '0'), ('22', '21', '22', '1', 'CaveJohson', '1593506839989', '0'), ('23', '21', '23', '1', 'CaveJohson', '1593507479378', '0'), ('24', '21', '24', '1', 'CaveJohson', '1593507499325', '0'), ('25', '21', '25', '1', 'CaveJohson', '1593507764518', '0'), ('26', '21', '26', '1', 'CaveJohson', '1593507791163', '0'), ('27', '21', '27', '1', 'CaveJohson', '1593507993386', '0');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
