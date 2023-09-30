<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet
    version="1.0"
    xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
    xmlns:svg="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink"
>

  <xsl:output method="xml" encoding="utf-8" indent="yes" omit-xml-declaration="yes" />

  <!-- Duplicate entire document -->
  <xsl:template match="node()|@*">
    <xsl:copy>
      <xsl:apply-templates select="node()|@*"/>
    </xsl:copy>
  </xsl:template>

  <!-- Reroute all anchor href attributes pointing at PlantUML files to SVG files instead -->
  <xsl:template match="svg:a/@xlink:href[substring(.,string-length()-4)='.puml']">
    <xsl:attribute name="xlink:href">
      <xsl:value-of select="concat(substring(.,1, string-length()-4), 'svg')"/>
    </xsl:attribute>
    <xsl:attribute name="href">
      <xsl:value-of select="concat(substring(.,1, string-length()-4), 'svg')"/>
    </xsl:attribute>
  </xsl:template>

  <!-- Reword all anchor title attributes referencing PlantUML files to instead reference SVG files -->
  <xsl:template match="svg:a/@xlink:title[substring(.,string-length()-4)='.puml']">
    <xsl:attribute name="xlink:title">
      <xsl:value-of select="concat(substring(.,1, string-length()-4), 'svg')"/>
    </xsl:attribute>
    <xsl:attribute name="title">
      <xsl:value-of select="concat(substring(.,1, string-length()-4), 'svg')"/>
    </xsl:attribute>
  </xsl:template>
</xsl:stylesheet>
