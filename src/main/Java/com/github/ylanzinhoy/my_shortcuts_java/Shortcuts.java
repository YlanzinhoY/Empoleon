package com.github.ylanzinhoy.my_shortcuts_java;

import com.intellij.codeInsight.template.TemplateActionContext;
import com.intellij.codeInsight.template.TemplateContextType;
import org.jetbrains.annotations.NotNull;

public class Shortcuts extends TemplateContextType {
    public Shortcuts() {
        super("SHORTCUTS_CONTEXT", "Shortcuts");
    }

    @Override
    public boolean isInContext(@NotNull TemplateActionContext templateActionContext) {
        return templateActionContext.getFile().getName().endsWith(".java");
    }


}
