package com.osigram.apz2024.mobile.ui

import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.sharp.Settings
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.style.TextOverflow

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TopBar(text: String, onClick: () -> Unit, modifier: Modifier = Modifier){
    TopAppBar(
        title = { Text(text, maxLines=1, overflow=TextOverflow.Ellipsis) },
        actions={
            IconButton(onClick = { onClick() }) {
                Icon(
                    Icons.Sharp.Settings,
                    null
                )
            }
        },
        modifier = modifier
    )
}